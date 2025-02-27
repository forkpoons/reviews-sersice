package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/forkpoons/reviews-sersice/internal/dto"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"time"
)

func (r *Review) GetQuestionsByStatus(ctx context.Context, productID uuid.UUID, status []string) (*[]dto.Question, error) {
	q := `SELECT * FROM reviews WHERE product_id = $1 AND rtype = 'question' AND status IN `
	if len(status) < 0 {
		return nil, errors.New(fmt.Sprintf("status is empty"))
	}
	statusStr := "("
	for _, str := range status {
		statusStr += `'` + str + `',`
	}
	statusStr = statusStr[:len(statusStr)-1] + `)`
	q += statusStr
	rows, err := r.pool.Query(ctx, q, productID)
	if err != nil {
		return nil, err
	}
	var questions []dto.Question
	reviews, err := pgx.CollectRows(rows, pgx.RowToStructByName[dto.ReviewDB])
	for _, review := range reviews {
		q = `SELECT * FROM answers WHERE question_id = $1 AND status IN ` + statusStr
		answerRows, err := r.pool.Query(ctx, q, review.ID)
		if err != nil {
			return nil, err
		}
		answers, err := pgx.CollectRows(answerRows, pgx.RowToStructByName[dto.Answer])
		questions = append(questions, dto.Question{
			ID:           review.ID,
			CreatedAt:    review.CreatedAt,
			UpdatedAt:    review.UpdatedAt,
			ProductId:    review.ProductId,
			UserID:       review.UserID,
			QuestionText: review.ReviewText,
			Media:        review.Media,
			Status:       review.Status,
			Answers:      answers,
		})
	}
	return &questions, nil
}

func (r *Review) AddQuestion(ctx context.Context, review dto.Question) error {
	q := `INSERT INTO reviews (id, rtype, created_at, updated_at, product_id, user_id, review_text, media, status) VALUES (@id, @type, @created_at, @updated_at, @product_id, @user_id, @review_text, @media, @status)`
	args := pgx.NamedArgs{
		"id":          uuid.New(),
		"type":        "question",
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"product_id":  review.ProductId,
		"user_id":     review.UserID,
		"review_text": review.QuestionText,
		"media":       review.Media,
		"status":      "created",
	}
	_, err := r.pool.Exec(ctx, q, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}
	return nil
}

func (r *Review) EditQuestion(ctx context.Context, review dto.Question) error {
	q := `UPDATE reviews SET review_text = @review_text, media = @media WHERE id = @id;`
	args := pgx.NamedArgs{
		"id":          review.ID,
		"review_text": review.QuestionText,
		"media":       review.Media,
	}
	_, err := r.pool.Exec(ctx, q, args)
	if err != nil {
		return err
	}
	return nil
}

func (r *Review) AddAnswer(ctx context.Context, answer dto.Answer) error {
	q := `INSERT INTO answers (id, answer_text, created_at, updated_at, question_id, user_id, status) VALUES (@id, @answer_text, @created_at, @updated_at, @question_id, @user_id, @status)`
	args := pgx.NamedArgs{
		"id":          uuid.New(),
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"question_id": answer.QuestionID,
		"user_id":     answer.UserID,
		"answer_text": answer.AnswerText,
		"status":      "created",
	}
	_, err := r.pool.Exec(ctx, q, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}
	return nil
}

func (r *Review) SetAnswerStatus(ctx context.Context, id uuid.UUID, status string) error {
	q := `UPDATE answers SET status = $1 WHERE id = $2`

	_, err := r.pool.Exec(ctx, q, status, id)
	if err != nil {
		return err
	}
	return nil
}
