type User struct {
  gorm.Model
  Profile   Profile
  ProfileID int
}

type Profile struct {
  gorm.Model
  Name string
}

db.Model(&user).Related(&profile)
//// SELECT * FROM profiles WHERE id = 111; // 111 is user's foreign key ProfileID
Specify Foreign Key

type Profile struct {
    gorm.Model
    Name string
}

type User struct {
    gorm.Model
    Profile      Profile `gorm:"ForeignKey:ProfileRefer"` // use ProfileRefer as foreign key
    ProfileRefer int
}
Specify Foreign Key & Association Key

type Profile struct {
    gorm.Model
    Refer string
    Name  string
}

type User struct {
    gorm.Model
    Profile   Profile `gorm:"ForeignKey:ProfileID;AssociationForeignKey:Refer"`
    ProfileID int
}
Has One

// User has one CreditCard, UserID is the foreign key
type User struct {
    gorm.Model
    CreditCard   CreditCard
}

type CreditCard struct {
    gorm.Model
    UserID   uint
    Number   string
}

var card CreditCard
db.Model(&user).Related(&card, "CreditCard")
//// SELECT * FROM credit_cards WHERE user_id = 123; // 123 is user's primary key
// CreditCard is user's field name, it means get user's CreditCard relations and fill it into variable card
// If the field name is same as the variable's type name, like above example, it could be omitted, like:
db.Model(&user).Related(&card)
Specify Foreign Key

type Profile struct {
  gorm.Model
  Name      string
  UserRefer uint
}

type User struct {
  gorm.Model
  Profile Profile `gorm:"ForeignKey:UserRefer"`
}
Specify Foreign Key & Association Key

type Profile struct {
  gorm.Model
  Name   string
  UserID uint
}

type User struct {
  gorm.Model
  Refer   string
  Profile Profile `gorm:"ForeignKey:UserID;AssociationForeignKey:Refer"`
}
Has Many

// User has many emails, UserID is the foreign key
type User struct {
    gorm.Model
    Emails   []Email
}

type Email struct {
    gorm.Model
    Email   string
    UserID  uint
}

db.Model(&user).Related(&emails)
//// SELECT * FROM emails WHERE user_id = 111; // 111 is user's primary key
Specify Foreign Key

type Profile struct {
  gorm.Model
  Name      string
  UserRefer uint
}

type User struct {
  gorm.Model
  Profiles []Profile `gorm:"ForeignKey:UserRefer"`
}
Specify Foreign Key & Association Key

type Profile struct {
  gorm.Model
  Name   string
  UserID uint
}

type User struct {
  gorm.Model
  Refer   string
  Profiles []Profile `gorm:"ForeignKey:UserID;AssociationForeignKey:Refer"`
}
Many To Many

// User has and belongs to many languages, use `user_languages` as join table
type User struct {
    gorm.Model
    Languages         []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
    gorm.Model
    Name string
}

db.Model(&user).Related(&languages)
//// SELECT * FROM "languages" INNER JOIN "user_languages" ON "user_languages"."language_id" = "languages"."id" WHERE "user_languages"."user_id" = 111
Specify Foreign Key & Association Key

type CustomizePerson struct {
  IdPerson string             `gorm:"primary_key:true"`
  Accounts []CustomizeAccount `gorm:"many2many:PersonAccount;AssociationForeignKey:idAccount;ForeignKey:idPerson"`
}

type CustomizeAccount struct {
  IdAccount string `gorm:"primary_key:true"`
  Name      string
}
Polymorphism

Supports polymorphic has-many and has-one associations.

  type Cat struct {
    Id    int
    Name  string
    Toy   Toy `gorm:"polymorphic:Owner;"`
  }

  type Dog struct {
    Id   int
    Name string
    Toy  Toy `gorm:"polymorphic:Owner;"`
  }

  type Toy struct {
    Id        int
    Name      string
    OwnerId   int
    OwnerType string
  }
Note: polymorphic belongs-to and many-to-many are explicitly NOT supported, and will throw errors.

Association Mode

Association Mode contains some helper methods to handle relationship things easily.

// Start Association Mode
var user User
db.Model(&user).Association("Languages")
// `user` is the source, it need to be a valid record (contains primary key)
// `Languages` is source's field name for a relationship.
// If those conditions not matched, will return an error, check it with:
// db.Model(&user).Association("Languages").Error


// Query - Find out all related associations
db.Model(&user).Association("Languages").Find(&languages)


// Append - Append new associations for many2many, has_many, will replace current association for has_one, belongs_to
db.Model(&user).Association("Languages").Append([]Language{languageZH, languageEN})
db.Model(&user).Association("Languages").Append(Language{Name: "DE"})


// Delete - Remove relationship between source & passed arguments, won't delete those arguments
db.Model(&user).Association("Languages").Delete([]Language{languageZH, languageEN})
db.Model(&user).Association("Languages").Delete(languageZH, languageEN)


// Replace - Replace current associations with new one
db.Model(&user).Association("Languages").Replace([]Language{languageZH, languageEN})
db.Model(&user).Association("Languages").Replace(Language{Name: "DE"}, languageEN)


// Count - Return the count of current associations
db.Model(&user).Association("Languages").Count()


// Clear - Remove relationship between source & current associations, won't delete those associations
db.Model(&user).Association("Languages").Clear()
