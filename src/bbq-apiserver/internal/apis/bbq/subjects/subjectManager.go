package subjects

import(
	"database/sql"
)

// IngredientType is an organization or categorization for subjects. 
type IngredientType struct {
	ID  			int  			`json:"id"`
	Name  			string  		`json:"name"`
	URLKey  		string  		`json:"urlkey"`
	Description  	string  		`json:"description"`
}

// Subject represents what is being bbqd 
type Subject struct {
	ID  			int  			`json:"id"`
	Name  			string  		`json:"name"`
	Description 	string  		`json:"description"`
	TypeID  		int  			`json:"typeid"`
	Type			IngredientType	`json:"type"`
}

// GetSubjectByID returns a subject structure given its ID
func GetSubjectByID(db *sql.DB, subjectID int) (Subject, error) {
	var s Subject
	
	if err := db.QueryRow("select id, name, description, typeid from bbq.subjects where id = $1", subjectID).Scan(&s.ID, &s.Name, &s.Description, &s.TypeID); err != nil {
		return Subject{}, err
	} 

	ingredientType, err := getIngredientTypeByID(db, s.TypeID);
	
	if err != nil {
		return Subject{}, err
	}

	s.Type = ingredientType

	return s, nil;
}

// GetSubjectByName returns a subject structure given its name
func GetSubjectByName(db *sql.DB, subjectName string)  (Subject, error) {
	var s Subject
	
	if err := db.QueryRow("select id, name, description, typeid from bbq.subjects where name = $1", subjectName).Scan(&s.ID, &s.Name, &s.Description, &s.TypeID); err != nil {
		return Subject{}, err
	} 

	ingredientType, err := getIngredientTypeByID(db, s.TypeID);
	
	if err != nil {
		return Subject{}, err
	}

	s.Type = ingredientType


	return s, nil;
}

func getIngredientTypeByID(db *sql.DB, ingredientTypeID int) (IngredientType, error) {
	var ingredientType IngredientType

	if err := db.QueryRow("select * from bbq.ingredient_types where id = $1", ingredientTypeID).Scan(&ingredientType.ID, &ingredientType.Name, &ingredientType.URLKey, &ingredientType.Description); err != nil {
		return IngredientType{}, err
	}

	return ingredientType, nil
}