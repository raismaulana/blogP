package apperror

/*
	ER400 Bad Request because request body is not satisfied (validation error)
	ER409 Conflict because resource already exists (validation error)
*/

const (
	FailUnmarshalResponseBodyError     ErrorType = "ER1000 Fail to unmarshal response body"        // used by controller
	ObjectNotFound                     ErrorType = "ER1001 Object %s is not found"                 // used by injected repo in interactor
	UnrecognizedEnum                   ErrorType = "ER1002 %s is not recognized %s enum"           // used by enum
	DatabaseNotFoundInContextError     ErrorType = "ER1003 Database is not found in context"       // used by repoimpl
	ValidationError                    ErrorType = "ER400 %s"                                      // used by repoimpl
	UsernameMustNotEmpty               ErrorType = "ER400 username must not empty"                 // used by entity/user
	NameMustNotEmpty                   ErrorType = "ER400 name must not empty"                     // used by entity/user
	EmailMustNotEmpty                  ErrorType = "ER400 email must not empty"                    // used by entity/user
	PasswordMustNotEmpty               ErrorType = "ER400 password must not empty"                 // used by entity/user
	CityMustNotEmpty                   ErrorType = "ER400 city must not empty"                     // used by entity/user
	CountryMustNotEmpty                ErrorType = "ER400 country must not empty"                  // used by entity/user
	BirthdayMustNotEmpty               ErrorType = "ER400 birthday must not empty"                 // used by entity/user
	WebProfileMustNotEmpty             ErrorType = "ER400 web profile must not empty"              // used by entity/user
	EmailAlreadyUsed                   ErrorType = "ER400 email already used"                      // used by entity/user
	UsernameAlreadyUsed                ErrorType = "ER400 username already used"                   // used by entity/user
	IdentifierMustNotEmpty             ErrorType = "ER400 identifier must not empty"               //
	NumberOnlyParam                    ErrorType = "ER400 number only param"                       //
	UserIsAlreadyActivated             ErrorType = "ER400 user is already activated"               //
	Unknown                            ErrorType = "ER400 unknown"                                 //
	InvalidActivationCode              ErrorType = "ER400 invalid activation code"                 //
	InvalidEmail                       ErrorType = "ER400 invalid email"                           //
	ActivationCodeIsIncorrectOrExpired ErrorType = "ER400 activation code is incorrect or expired" //
	InvalidCredential                  ErrorType = "ER400 invalid credential"                      //
	FailedGenerateAuthToken            ErrorType = "ER400 failed generate auth token"              //
	TagMustNotEmpty                    ErrorType = "ER400 tag must not empty"                      //
	TagAlreadyExsist                   ErrorType = "ER400 tag already exsist"                      //
	CategoryMustNotEmpty               ErrorType = "ER400 category must not empty"                 //
	CategoryAlreadyExsist              ErrorType = "ER400 category already exsist"                 //
	DescriptionMustNotEmpty            ErrorType = "ER1000 description must not empty"             //
	TitleMustNotEmpty                  ErrorType = "ER1000 title must not empty"                   //
	ContentMustBeValidJSON             ErrorType = "ER1000 content must be valid json"             //
	AuthorIDMustNotEmpty               ErrorType = "ER1000 author id must not empty"               //
	SlugMustNotEmpty                   ErrorType = "ER1000 slug must not empty"                    //
	CoverMustNotEmpty                  ErrorType = "ER1000 cover must not empty"                   //
	InvalidToken                       ErrorType = "ER1000 invalid token"                          //
	SlugAlreadyExsist                  ErrorType = "ER1000 slug already exsist"                    //
	SomeCategoryDoesNotExist           ErrorType = "ER1000 some category does not exist"           //
	SomeTagDoesNotExist                ErrorType = "ER1000 some tag does not exist"                //
	Forbidden                          ErrorType = "ER1000 forbidden"                              //
)
