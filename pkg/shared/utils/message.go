package utils

// ***********Database***********
// Error messages
const (
	DatabaseConnectionError = "error connecting to the database"
	UnexpectedDatabaseError = "unexpected database error"
	SchemaMigrationSuccess  = "%s schema migrated successfully"
	DatabaseMigrationError  = "failed to migrate %s schema: %v"
)

// *********** General Errors ***********
const (
	InternalServerError = "an unexpected error occurred"
	NotFoundError       = "resource not found"
	BadRequestError     = "bad request"
	UnauthorizedError   = "unauthorized access"
	ForbiddenError      = "access forbidden"
)

// ***********Product***********
// Error messages
const (
	ProductNotFoundError         = "product with ID %d not found"
	ProductDeletionError         = "error occurred while deleting product with ID %d"
	ProductUpdateError           = "error occurred while updating product with ID %d"
	ProductCreationError         = "error occurred while creating product"
	InvalidProductIDError        = "invalid product ID provided"
	InvalidProductDataError      = "invalid product data"
	InvalidProductRequest        = "Invalid product request"
	ProductUnexpectedFetchError  = "unexpected error fetching product for update"
	ProductUnexpectedUpdateError = "unexpected error while fetching product"
)

// Info messages
const (
	ProductsFetchedSuccessfully = "Fetched all products successfully"
	ProductFetchedSuccessfully  = "Fetched product successfully with ID %d"
	ProductCreatedSuccessfully  = "product created successfully with ID %d"
	ProductUpdatedSuccessfully  = "product updated successfully with ID %d"
	ProductDeletedSuccessfully  = "product deleted successfully with ID %d"
)

// Validation error messages
const (
	InvalidRequestMethod = "invalid request method"
	InvalidRequestPath   = "invalid request path"
	InvalidRequestBody   = "invalid request body"
)
