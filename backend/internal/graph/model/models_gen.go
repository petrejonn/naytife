// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type CreateCategoryAttributePayload interface {
	IsCreateCategoryAttributePayload()
}

type CreateCategoryPayload interface {
	IsCreateCategoryPayload()
}

type CreateProductPayload interface {
	IsCreateProductPayload()
}

type CreateShopPayload interface {
	IsCreateShopPayload()
}

type CreateWhatsAppPayload interface {
	IsCreateWhatsAppPayload()
}

type DeleteCategoryAttributePayload interface {
	IsDeleteCategoryAttributePayload()
}

type Node interface {
	IsNode()
	GetID() string
}

type SignInUserPayload interface {
	IsSignInUserPayload()
}

type SocialMediaContact interface {
	IsSocialMediaContact()
	GetURL() *string
}

type UpdateCategoryPayload interface {
	IsUpdateCategoryPayload()
}

type UpdateShopPayload interface {
	IsUpdateShopPayload()
}

type UpdateWhatsAppPayload interface {
	IsUpdateWhatsAppPayload()
}

type UserError interface {
	IsUserError()
	GetMessage() string
	GetCode() ErrorCode
	GetPath() []string
}

type AllowedCategoryAttributes struct {
	Title    string                   `json:"title"`
	DataType ProductAttributeDataType `json:"dataType"`
}

type AllowedProductAttributes struct {
	Title    string                   `json:"title"`
	DataType ProductAttributeDataType `json:"dataType"`
}

type Category struct {
	ID                string                      `json:"id"`
	Slug              string                      `json:"slug"`
	Title             string                      `json:"title"`
	Description       string                      `json:"description"`
	Parent            *Category                   `json:"parent,omitempty"`
	Children          []Category                  `json:"children,omitempty"`
	Products          *ProductConnection          `json:"products,omitempty"`
	AllowedAttributes []AllowedCategoryAttributes `json:"allowedAttributes"`
	Images            *CategoryImages             `json:"images,omitempty"`
	UpdatedAt         time.Time                   `json:"updatedAt"`
	CreatedAt         time.Time                   `json:"createdAt"`
}

func (Category) IsNode()            {}
func (this Category) GetID() string { return this.ID }

type CategoryConnection struct {
	Edges    []CategoryEdge `json:"edges"`
	PageInfo *PageInfo      `json:"pageInfo"`
}

type CategoryEdge struct {
	Cursor string    `json:"cursor"`
	Node   *Category `json:"node"`
}

type CategoryImages struct {
	Banner *Image `json:"banner"`
}

type CategoryNotFoundError struct {
	Message string    `json:"message"`
	Code    ErrorCode `json:"code"`
	Path    []string  `json:"path"`
}

func (CategoryNotFoundError) IsUserError()            {}
func (this CategoryNotFoundError) GetMessage() string { return this.Message }
func (this CategoryNotFoundError) GetCode() ErrorCode { return this.Code }
func (this CategoryNotFoundError) GetPath() []string {
	if this.Path == nil {
		return nil
	}
	interfaceSlice := make([]string, 0, len(this.Path))
	for _, concrete := range this.Path {
		interfaceSlice = append(interfaceSlice, concrete)
	}
	return interfaceSlice
}

func (CategoryNotFoundError) IsUpdateCategoryPayload() {}

func (CategoryNotFoundError) IsCreateCategoryAttributePayload() {}

func (CategoryNotFoundError) IsDeleteCategoryAttributePayload() {}

func (CategoryNotFoundError) IsCreateProductPayload() {}

type CreateCategoryAttributeInput struct {
	Title    string                   `json:"title"`
	DataType ProductAttributeDataType `json:"dataType"`
}

type CreateCategoryAttributeSuccess struct {
	Attributes []AllowedCategoryAttributes `json:"attributes"`
}

func (CreateCategoryAttributeSuccess) IsCreateCategoryAttributePayload() {}

type CreateCategoryInput struct {
	ParentID    *string `json:"parentID,omitempty"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
}

type CreateCategorySuccess struct {
	Category *Category `json:"category,omitempty"`
}

func (CreateCategorySuccess) IsCreateCategoryPayload() {}

type CreateProductInput struct {
	CategoryID  string  `json:"categoryID"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type CreateProductSuccess struct {
	Product *Product `json:"product"`
}

func (CreateProductSuccess) IsCreateProductPayload() {}

type CreateShopInput struct {
	Title  string `json:"title"`
	Domain string `json:"domain"`
}

type CreateShopSuccess struct {
	Shop *Shop `json:"shop,omitempty"`
}

func (CreateShopSuccess) IsCreateShopPayload() {}

type CreateWhatsAppInput struct {
	URL         string            `json:"url"`
	PhoneNumber *PhoneNumberInput `json:"phoneNumber"`
}

type CreateWhatsAppSuccess struct {
	WhatsApp *WhatsApp `json:"whatsApp,omitempty"`
}

func (CreateWhatsAppSuccess) IsCreateWhatsAppPayload() {}

type DeleteCategoryAttributeSuccess struct {
	Attributes []AllowedCategoryAttributes `json:"attributes"`
}

func (DeleteCategoryAttributeSuccess) IsDeleteCategoryAttributePayload() {}

type Facebook struct {
	URL    string `json:"url"`
	Handle string `json:"handle"`
}

func (Facebook) IsSocialMediaContact() {}
func (this Facebook) GetURL() *string  { return &this.URL }

type Image struct {
	URL     string  `json:"url"`
	AltText *string `json:"altText,omitempty"`
}

type Instagram struct {
	URL    string `json:"url"`
	Handle string `json:"handle"`
}

func (Instagram) IsSocialMediaContact() {}
func (this Instagram) GetURL() *string  { return &this.URL }

type Mutation struct {
}

type PageInfo struct {
	EndCursor       *string `json:"endCursor,omitempty"`
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor,omitempty"`
}

type PhoneNumber struct {
	E164 string `json:"e164"`
}

type PhoneNumberInput struct {
	E164 string `json:"e164"`
}

type Product struct {
	ID                string                     `json:"id"`
	Slug              string                     `json:"slug"`
	Title             string                     `json:"title"`
	Price             float64                    `json:"price"`
	Description       string                     `json:"description"`
	Category          *Category                  `json:"category"`
	DefaultVariant    *ProductVariant            `json:"defaultVariant"`
	Variants          []ProductVariant           `json:"variants"`
	AllowedAttributes []AllowedProductAttributes `json:"allowedAttributes"`
	Images            []Image                    `json:"images"`
	Status            *ProductStatus             `json:"status,omitempty"`
	UpdatedAt         time.Time                  `json:"updatedAt"`
	CreatedAt         time.Time                  `json:"createdAt"`
}

func (Product) IsNode()            {}
func (this Product) GetID() string { return this.ID }

type ProductAttribute struct {
	Key   string  `json:"key"`
	Value *string `json:"value,omitempty"`
}

type ProductAttributeValue struct {
	IntValue    *int    `json:"intValue,omitempty"`
	StringValue *string `json:"stringValue,omitempty"`
}

type ProductConnection struct {
	Edges    []ProductEdge `json:"edges"`
	PageInfo *PageInfo     `json:"pageInfo"`
}

type ProductEdge struct {
	Cursor string   `json:"cursor"`
	Node   *Product `json:"node"`
}

type ProductVariant struct {
	ID                string             `json:"id"`
	Slug              string             `json:"slug"`
	Title             string             `json:"title"`
	Price             float64            `json:"price"`
	AvailableQuantity int                `json:"availableQuantity"`
	Description       string             `json:"description"`
	Attributes        []ProductAttribute `json:"attributes"`
	StockStatus       ProductStockStatus `json:"stockStatus"`
	UpdatedAt         time.Time          `json:"updatedAt"`
	CreatedAt         time.Time          `json:"createdAt"`
}

func (ProductVariant) IsNode()            {}
func (this ProductVariant) GetID() string { return this.ID }

type Query struct {
}

type Shop struct {
	ID                   string              `json:"id"`
	Title                string              `json:"title"`
	DefaultDomain        string              `json:"defaultDomain"`
	ContactPhone         *PhoneNumber        `json:"contactPhone,omitempty"`
	ContactEmail         *string             `json:"contactEmail,omitempty"`
	Address              *ShopAddress        `json:"address,omitempty"`
	Products             *ProductConnection  `json:"products,omitempty"`
	Categories           *CategoryConnection `json:"categories,omitempty"`
	WhatsApp             *WhatsApp           `json:"whatsApp,omitempty"`
	Facebook             *Facebook           `json:"facebook,omitempty"`
	Images               *ShopImages         `json:"images,omitempty"`
	CurrencyCode         string              `json:"currencyCode"`
	Status               ShopStatus          `json:"status"`
	About                *string             `json:"about,omitempty"`
	ShopProductsCategory *string             `json:"shopProductsCategory,omitempty"`
	SeoDescription       *string             `json:"seoDescription,omitempty"`
	SeoKeywords          []string            `json:"seoKeywords"`
	SeoTitle             *string             `json:"seoTitle,omitempty"`
	UpdatedAt            time.Time           `json:"updatedAt"`
	CreatedAt            time.Time           `json:"createdAt"`
	Owner                *User               `json:"owner"`
}

func (Shop) IsNode()            {}
func (this Shop) GetID() string { return this.ID }

type ShopAddress struct {
	Address string `json:"address"`
}

type ShopAddressInput struct {
	Address string `json:"address"`
}

type ShopImages struct {
	SiteLogo   *Image `json:"siteLogo,omitempty"`
	Favicon    *Image `json:"favicon,omitempty"`
	Banner     *Image `json:"banner,omitempty"`
	CoverImage *Image `json:"coverImage,omitempty"`
}

type ShopNotFoundError struct {
	Message string    `json:"message"`
	Code    ErrorCode `json:"code"`
	Path    []string  `json:"path"`
}

func (ShopNotFoundError) IsUserError()            {}
func (this ShopNotFoundError) GetMessage() string { return this.Message }
func (this ShopNotFoundError) GetCode() ErrorCode { return this.Code }
func (this ShopNotFoundError) GetPath() []string {
	if this.Path == nil {
		return nil
	}
	interfaceSlice := make([]string, 0, len(this.Path))
	for _, concrete := range this.Path {
		interfaceSlice = append(interfaceSlice, concrete)
	}
	return interfaceSlice
}

func (ShopNotFoundError) IsUpdateShopPayload() {}

func (ShopNotFoundError) IsCreateWhatsAppPayload() {}

func (ShopNotFoundError) IsUpdateWhatsAppPayload() {}

type SignInInput struct {
	Username *string `json:"username,omitempty"`
}

type SignInUserSuccess struct {
	User *User `json:"user,omitempty"`
}

func (SignInUserSuccess) IsSignInUserPayload() {}

type UpdateCategoryInput struct {
	ParentID    *string `json:"parentID,omitempty"`
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
}

type UpdateCategorySuccess struct {
	Category *Category `json:"category"`
}

func (UpdateCategorySuccess) IsUpdateCategoryPayload() {}

type UpdateShopInput struct {
	Title          *string           `json:"title,omitempty"`
	ContactEmail   *string           `json:"contactEmail,omitempty"`
	ContactPhone   *PhoneNumberInput `json:"contactPhone,omitempty"`
	Address        *ShopAddressInput `json:"address,omitempty"`
	CurrencyCode   *string           `json:"currencyCode,omitempty"`
	About          *string           `json:"about,omitempty"`
	SeoDescription *string           `json:"seoDescription,omitempty"`
	SeoKeywords    []string          `json:"seoKeywords,omitempty"`
	SeoTitle       *string           `json:"seoTitle,omitempty"`
}

type UpdateShopSuccess struct {
	Shop *Shop `json:"shop,omitempty"`
}

func (UpdateShopSuccess) IsUpdateShopPayload() {}

type UpdateWhatsAppInput struct {
	URL         *string           `json:"url,omitempty"`
	PhoneNumber *PhoneNumberInput `json:"phoneNumber,omitempty"`
}

type UpdateWhatsAppSuccess struct {
	WhatsApp *WhatsApp `json:"whatsApp"`
}

func (UpdateWhatsAppSuccess) IsUpdateWhatsAppPayload() {}

type User struct {
	ID                string  `json:"id"`
	Email             string  `json:"email"`
	Name              *string `json:"name,omitempty"`
	ProfilePictureURL *string `json:"profilePictureUrl,omitempty"`
	CreatedAt         string  `json:"createdAt"`
	LastLogin         string  `json:"lastLogin"`
}

func (User) IsNode()            {}
func (this User) GetID() string { return this.ID }

type WhatsApp struct {
	URL         string       `json:"url"`
	PhoneNumber *PhoneNumber `json:"phoneNumber"`
}

func (WhatsApp) IsSocialMediaContact() {}
func (this WhatsApp) GetURL() *string  { return &this.URL }

type ErrorCode string

const (
	ErrorCodeNotFoundShop           ErrorCode = "NOT_FOUND_SHOP"
	ErrorCodeNotFoundCategory       ErrorCode = "NOT_FOUND_CATEGORY"
	ErrorCodeAuthInvalidToken       ErrorCode = "AUTH_INVALID_TOKEN"
	ErrorCodeValidationInvalidInput ErrorCode = "VALIDATION_INVALID_INPUT"
	ErrorCodeServerErrorInternal    ErrorCode = "SERVER_ERROR_INTERNAL"
	ErrorCodeRateLimitExceeded      ErrorCode = "RATE_LIMIT_EXCEEDED"
)

var AllErrorCode = []ErrorCode{
	ErrorCodeNotFoundShop,
	ErrorCodeNotFoundCategory,
	ErrorCodeAuthInvalidToken,
	ErrorCodeValidationInvalidInput,
	ErrorCodeServerErrorInternal,
	ErrorCodeRateLimitExceeded,
}

func (e ErrorCode) IsValid() bool {
	switch e {
	case ErrorCodeNotFoundShop, ErrorCodeNotFoundCategory, ErrorCodeAuthInvalidToken, ErrorCodeValidationInvalidInput, ErrorCodeServerErrorInternal, ErrorCodeRateLimitExceeded:
		return true
	}
	return false
}

func (e ErrorCode) String() string {
	return string(e)
}

func (e *ErrorCode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ErrorCode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ErrorCode", str)
	}
	return nil
}

func (e ErrorCode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ProductAttributeDataType string

const (
	ProductAttributeDataTypeString  ProductAttributeDataType = "STRING"
	ProductAttributeDataTypeInteger ProductAttributeDataType = "INTEGER"
)

var AllProductAttributeDataType = []ProductAttributeDataType{
	ProductAttributeDataTypeString,
	ProductAttributeDataTypeInteger,
}

func (e ProductAttributeDataType) IsValid() bool {
	switch e {
	case ProductAttributeDataTypeString, ProductAttributeDataTypeInteger:
		return true
	}
	return false
}

func (e ProductAttributeDataType) String() string {
	return string(e)
}

func (e *ProductAttributeDataType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProductAttributeDataType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProductAttributeDataType", str)
	}
	return nil
}

func (e ProductAttributeDataType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ProductStatus string

const (
	ProductStatusDraft     ProductStatus = "DRAFT"
	ProductStatusPublished ProductStatus = "PUBLISHED"
	ProductStatusArchived  ProductStatus = "ARCHIVED"
)

var AllProductStatus = []ProductStatus{
	ProductStatusDraft,
	ProductStatusPublished,
	ProductStatusArchived,
}

func (e ProductStatus) IsValid() bool {
	switch e {
	case ProductStatusDraft, ProductStatusPublished, ProductStatusArchived:
		return true
	}
	return false
}

func (e ProductStatus) String() string {
	return string(e)
}

func (e *ProductStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProductStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProductStatus", str)
	}
	return nil
}

func (e ProductStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ProductStockStatus string

const (
	ProductStockStatusInStock    ProductStockStatus = "IN_STOCK"
	ProductStockStatusOutOfStock ProductStockStatus = "OUT_OF_STOCK"
	ProductStockStatusPreorder   ProductStockStatus = "PREORDER"
)

var AllProductStockStatus = []ProductStockStatus{
	ProductStockStatusInStock,
	ProductStockStatusOutOfStock,
	ProductStockStatusPreorder,
}

func (e ProductStockStatus) IsValid() bool {
	switch e {
	case ProductStockStatusInStock, ProductStockStatusOutOfStock, ProductStockStatusPreorder:
		return true
	}
	return false
}

func (e ProductStockStatus) String() string {
	return string(e)
}

func (e *ProductStockStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProductStockStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProductStockStatus", str)
	}
	return nil
}

func (e ProductStockStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ShopStatus string

const (
	ShopStatusDraft     ShopStatus = "DRAFT"
	ShopStatusPublished ShopStatus = "PUBLISHED"
	ShopStatusArchived  ShopStatus = "ARCHIVED"
	ShopStatusSuspended ShopStatus = "SUSPENDED"
)

var AllShopStatus = []ShopStatus{
	ShopStatusDraft,
	ShopStatusPublished,
	ShopStatusArchived,
	ShopStatusSuspended,
}

func (e ShopStatus) IsValid() bool {
	switch e {
	case ShopStatusDraft, ShopStatusPublished, ShopStatusArchived, ShopStatusSuspended:
		return true
	}
	return false
}

func (e ShopStatus) String() string {
	return string(e)
}

func (e *ShopStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ShopStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ShopStatus", str)
	}
	return nil
}

func (e ShopStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
