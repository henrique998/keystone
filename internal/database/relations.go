package database

import "github.com/henrique998/keystone/orm"

type RelationType int

const (
	RelationBelongsTo RelationType = iota + 1
)

type BelongsTo[T any] struct {
	localColumn  string
	refTable     string
	refColumn    string
	relationName string
}

type RelationMetadata struct {
	Type        RelationType
	LocalColumn string
	RefTable    string
	RefColumn   string
	Name        string
}

func NewBelongsTo[T any](
	localColumn string,
	refTable string,
	refColumn string,
	relationName string,
) BelongsTo[T] {
	return BelongsTo[T]{
		localColumn:  localColumn,
		refTable:     refTable,
		refColumn:    refColumn,
		relationName: relationName,
	}
}

func (b BelongsTo[T]) Metadata() RelationMetadata {
	return RelationMetadata{
		Type:        RelationBelongsTo,
		LocalColumn: b.localColumn,
		RefTable:    b.refTable,
		RefColumn:   b.refColumn,
		Name:        b.relationName,
	}
}

type withRelation struct {
	meta RelationMetadata
}

func (w withRelation) IsWithOption() {}

func WithRelation(meta RelationMetadata) orm.WithOption {
	return withRelation{meta: meta}
}
