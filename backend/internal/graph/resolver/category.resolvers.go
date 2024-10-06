package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.54

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/gosimple/slug"
	"github.com/petrejonn/naytife/internal/db"
	"github.com/petrejonn/naytife/internal/graph/generated"
	"github.com/petrejonn/naytife/internal/graph/model"
)

// ID is the resolver for the id field.
func (r *categoryResolver) ID(ctx context.Context, obj *model.Category) (string, error) {
	// Return the base64-encoded ID
	return encodeRelayID("Category", obj.ID), nil
}

// Children is the resolver for the children field.
func (r *categoryResolver) Children(ctx context.Context, obj *model.Category) ([]model.Category, error) {
	shopID := ctx.Value("shop_id").(int64)
	catID, err := strconv.ParseInt(obj.ID, 10, 64)
	if err != nil {
		return nil, errors.New("invalid category")
	}
	categoriesDB, err := r.Repository.GetCategoryChildren(ctx, db.GetCategoryChildrenParams{ShopID: shopID, ParentID: &catID})
	if err != nil {
		return nil, errors.New("could not fetch categories")
	}
	categories := make([]model.Category, 0, len(categoriesDB))
	for _, cat := range categoriesDB {
		// attributes, err := unmarshalCategoryAttributes(cat.CategoryAttributes)
		categories = append(categories, model.Category{
			ID:          strconv.FormatInt(cat.CategoryID, 10),
			Slug:        cat.Slug,
			Title:       cat.Title,
			Description: cat.Description,
			CreatedAt:   cat.CreatedAt.Time,
			UpdatedAt:   cat.UpdatedAt.Time,
		})
	}
	return categories, nil
}

// Products is the resolver for the products field.
func (r *categoryResolver) Products(ctx context.Context, obj *model.Category, first *int, after *string) (*model.ProductConnection, error) {
	catID, err := strconv.Atoi(obj.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid category id %v", obj.ID)
	}
	limit := 20
	if first != nil {
		limit = *first
	}
	afterID := int64(0)
	if after != nil {
		decodedType, id, err := decodeRelayID(*after)
		if err != nil {
			return nil, fmt.Errorf("invalid after cursor: %w", err)
		}
		if decodedType != "Product" {
			return nil, fmt.Errorf("expected after cursor type 'Product', got '%s'", decodedType)
		}
		if id != nil {
			afterID = *id
		}
	}
	productsDB, err := r.Repository.GetProductsByCategory(ctx, db.GetProductsByCategoryParams{CategoryID: int64(catID), After: afterID, Limit: int32(limit) + 1})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch products: %w", err)
	}
	hasNextPage := len(productsDB) > limit
	if hasNextPage {
		productsDB = productsDB[:limit]
	}
	edges := make([]model.ProductEdge, len(productsDB))
	for i, prod := range productsDB {
		relayID := encodeRelayID("Product", strconv.FormatInt(prod.ProductID, 10))
		edges[i] = model.ProductEdge{Cursor: relayID, Node: &model.Product{
			ID:          strconv.FormatInt(prod.ProductID, 10),
			Title:       prod.Title,
			Description: prod.Description,
			CreatedAt:   prod.CreatedAt.Time,
			UpdatedAt:   prod.UpdatedAt.Time,
			Status:      (*model.ProductStatus)(&prod.Status),
		}}
	}
	var startCursor, endCursor *string
	if len(productsDB) > 0 {
		firstCursor := encodeRelayID("Product", strconv.FormatInt(productsDB[0].ProductID, 10))
		lastCursor := encodeRelayID("Product", strconv.FormatInt(productsDB[len(productsDB)-1].ProductID, 10))
		startCursor, endCursor = &firstCursor, &lastCursor
	}

	pageInfo := &model.PageInfo{
		HasNextPage: hasNextPage,
		StartCursor: safeStringDereference(startCursor),
		EndCursor:   safeStringDereference(endCursor),
	}

	return &model.ProductConnection{
		Edges:      edges,
		PageInfo:   pageInfo,
		TotalCount: len(productsDB),
	}, nil
}

// AllowedAttributes is the resolver for the allowedAttributes field.
func (r *categoryResolver) AllowedAttributes(ctx context.Context, obj *model.Category) ([]model.AllowedCategoryAttributes, error) {
	catID, err := strconv.Atoi(obj.ID)
	if err != nil {
		return nil, errors.New("invalid category id")
	}
	attributesDB, err := r.Repository.GetCategoryAttributes(ctx, int64(catID))
	if err != nil {
		return nil, errors.New("could not fetch category attribute")
	}
	attributes, err := unmarshalCategoryAttributes(attributesDB)
	if err != nil {
		return nil, errors.New("failed to get attributes")
	}
	return attributes, nil
}

// Images is the resolver for the images field.
func (r *categoryResolver) Images(ctx context.Context, obj *model.Category) (*model.CategoryImages, error) {
	panic(fmt.Errorf("not implemented: Images - images"))
}

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, category model.CreateCategoryInput) (model.CreateCategoryPayload, error) {
	shopID := ctx.Value("shop_id").(int64)
	param := db.CreateCategoryParams{
		Title:       category.Title,
		Description: category.Description,
		Slug:        slug.MakeLang(category.Title, "en"),
		ShopID:      shopID,
	}
	param.CategoryAttributes = []byte("{}")
	if category.ParentID != nil {
		_, catID, err := decodeRelayID(*category.ParentID)
		if err != nil {
			return nil, errors.New("could not find parent category")
		}
		parent, err := r.Repository.GetCategory(ctx, db.GetCategoryParams{ShopID: shopID, CategoryID: *catID})
		if err != nil {
			return nil, errors.New("could not find parent category")
		}
		param.ParentID = &parent.CategoryID
		param.CategoryAttributes = parent.CategoryAttributes
	}
	cat, err := r.Repository.CreateCategory(ctx, param)
	if err != nil {
		return nil, err
	}
	return &model.CreateCategorySuccess{
		Category: &model.Category{
			ID:          strconv.FormatInt(cat.CategoryID, 10),
			Slug:        cat.Slug,
			Title:       cat.Title,
			Description: cat.Description,
			CreatedAt:   cat.CreatedAt.Time,
			UpdatedAt:   cat.UpdatedAt.Time,
		}}, nil
}

// UpdateCategory is the resolver for the updateCategory field.
func (r *mutationResolver) UpdateCategory(ctx context.Context, categoryID string, category model.UpdateCategoryInput) (model.UpdateCategoryPayload, error) {
	_, catId, err := decodeRelayID(categoryID)
	if err != nil {
		return nil, errors.New("invalid category ID")
	}
	params := db.UpdateCategoryParams{
		CategoryID:  *catId,
		Title:       category.Title,
		Description: category.Description,
	}

	if category.ParentID != nil {
		_, id, err := decodeRelayID(*category.ParentID)
		if err != nil {
			return nil, errors.New("could not find parent category")
		}
		params.ParentID = id
	}

	dbCat, err := r.Repository.UpdateCategory(ctx, params)
	if err != nil {
		return nil, errors.New("could not update category")
	}
	return &model.UpdateCategorySuccess{
		Category: &model.Category{
			ID:          strconv.FormatInt(dbCat.CategoryID, 10),
			Slug:        dbCat.Slug,
			Title:       dbCat.Title,
			Description: dbCat.Description,
			CreatedAt:   dbCat.CreatedAt.Time,
			UpdatedAt:   dbCat.UpdatedAt.Time,
		}}, nil
}

// CreateCategoryAttribute is the resolver for the createCategoryAttribute field.
func (r *mutationResolver) CreateCategoryAttribute(ctx context.Context, categoryID string, attribute model.CreateCategoryAttributeInput) (model.CreateCategoryAttributePayload, error) {
	_, catId, err := decodeRelayID(categoryID)
	if err != nil {
		return nil, errors.New("invalid category ID")
	}
	attributesDB, err := r.Repository.CreateCategoryAttribute(ctx, db.CreateCategoryAttributeParams{
		CategoryID: *catId,
		Title:      attribute.Title,
		DataType:   attribute.DataType.String(),
	})
	if err != nil {
		return nil, errors.New("could not create category attribute")
	}
	attributes, err := unmarshalCategoryAttributes(attributesDB)
	if err != nil {
		return nil, errors.New("failed to fetch attributes")
	}
	return &model.CreateCategoryAttributeSuccess{
		Attributes: attributes,
	}, nil
}

// DeleteCategoryAttribute is the resolver for the deleteCategoryAttribute field.
func (r *mutationResolver) DeleteCategoryAttribute(ctx context.Context, categoryID string, attribute string) (model.DeleteCategoryAttributePayload, error) {
	_, catId, err := decodeRelayID(categoryID)
	if err != nil {
		return nil, errors.New("invalid category ID")
	}
	attributesDB, err := r.Repository.DeleteCategoryAttribute(ctx, db.DeleteCategoryAttributeParams{
		CategoryID: *catId, Attribute: attribute,
	})
	if err != nil {
		return nil, errors.New("could not fetch category attribute")
	}
	attributes, err := unmarshalCategoryAttributes(attributesDB)
	if err != nil {
		return nil, errors.New("failed to fetch attributes")
	}
	return &model.DeleteCategoryAttributeSuccess{
		Attributes: attributes,
	}, nil
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context, first *int, after *string) (*model.CategoryConnection, error) {
	shopID := ctx.Value("shop_id").(int64)
	limit := 20
	if first != nil {
		limit = *first
	}
	afterID := int64(0)
	if after != nil {
		decodedType, id, err := decodeRelayID(*after)
		if err != nil {
			return nil, fmt.Errorf("invalid after cursor: %w", err)
		}
		if decodedType != "Category" {
			return nil, fmt.Errorf("expected after cursor type 'Category', got '%s'", decodedType)
		}
		if id != nil {
			afterID = *id
		}
	}
	categoriesDB, err := r.Repository.GetCategories(ctx, db.GetCategoriesParams{ShopID: shopID, After: afterID, Limit: int32(limit) + 1})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch categories: %w", err)
	}
	hasNextPage := len(categoriesDB) > limit
	if hasNextPage {
		categoriesDB = categoriesDB[:limit]
	}
	edges := make([]model.CategoryEdge, len(categoriesDB))
	for i, cat := range categoriesDB {
		relayID := encodeRelayID("Category", strconv.FormatInt(cat.CategoryID, 10))
		edges[i] = model.CategoryEdge{Cursor: relayID, Node: &model.Category{
			ID:          strconv.FormatInt(cat.CategoryID, 10),
			Slug:        cat.Slug,
			Title:       cat.Title,
			Description: cat.Description,
			CreatedAt:   cat.CreatedAt.Time,
			UpdatedAt:   cat.UpdatedAt.Time,
		}}
	}
	var startCursor, endCursor *string
	if len(categoriesDB) > 0 {
		firstCursor := encodeRelayID("Category", strconv.FormatInt(categoriesDB[0].CategoryID, 10))
		lastCursor := encodeRelayID("Category", strconv.FormatInt(categoriesDB[len(categoriesDB)-1].CategoryID, 10))
		startCursor, endCursor = &firstCursor, &lastCursor
	}

	pageInfo := &model.PageInfo{
		HasNextPage: hasNextPage,
		StartCursor: safeStringDereference(startCursor),
		EndCursor:   safeStringDereference(endCursor),
	}

	// Return the CategoryConnection result
	return &model.CategoryConnection{
		Edges:      edges,
		PageInfo:   pageInfo,
		TotalCount: len(categoriesDB),
	}, nil
}

// Category is the resolver for the category field.
func (r *queryResolver) Category(ctx context.Context, id string) (*model.Category, error) {
	shopID := ctx.Value("shop_id").(int64)
	_, catID, err := decodeRelayID(id)
	if err != nil {
		return nil, errors.New("invalid category ID")
	}
	cat, err := r.Repository.GetCategory(ctx, db.GetCategoryParams{ShopID: shopID, CategoryID: *catID})
	if err != nil {
		return nil, errors.New("could not find category")
	}
	return &model.Category{
		ID:          strconv.FormatInt(cat.CategoryID, 10),
		Slug:        cat.Slug,
		Title:       cat.Title,
		Description: cat.Description,
		CreatedAt:   cat.CreatedAt.Time,
		UpdatedAt:   cat.UpdatedAt.Time,
	}, nil
}

// Category returns generated.CategoryResolver implementation.
func (r *Resolver) Category() generated.CategoryResolver { return &categoryResolver{r} }

type categoryResolver struct{ *Resolver }
