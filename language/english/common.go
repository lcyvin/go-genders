package english

import (
  "github.com/lcyvin/go-genders"
)

var ZeHir genders.Pronoun = genders.Pronoun{
  Language: "english",
  Alias: "Ze/Hir",
  ThirdPerson: genders.PronounPart{
    Nominative: "ze",
    Dative: "hir",
    Genitive: "hir",
    GenitiveDeterminer: "hirs",
    Reflexive: "hirself",
  },
}

var FaeFaer genders.Pronoun = genders.Pronoun{
  Language: "english",
  Alias: "Fae/Faer",
  ThirdPerson: genders.PronounPart{
    Nominative: "fae",
    Dative: "faer",
    Genitive: "faer",
    GenitiveDeterminer: "faers",
    Reflexive: "faerself",
  },
}

var TheyThem genders.Pronoun = genders.Pronoun{
  Language: "english",
  Alias: "They/Them",
  ThirdPerson: genders.PronounPart{
    Nominative: "they",
    Dative: "them",
    Genitive: "theirs",
    GenitiveDeterminer: "their",
    Reflexive: "themself",
  },
}

var HeHim genders.Pronoun = genders.Pronoun{
  Language: "english",
  Alias: "He/Him",
  ThirdPerson: genders.PronounPart{
    Nominative: "he",
    Dative: "him",
    Genitive: "his",
    GenitiveDeterminer: "his",
    Reflexive: "himself",
  },
}

var SheHer genders.Pronoun = genders.Pronoun{
  Language: "english",
  Alias: "She/Her",
  ThirdPerson: genders.PronounPart{
    Nominative: "she",
    Dative: "her",
    Genitive: "hers",
    GenitiveDeterminer: "hers",
    Reflexive: "herself",
  },
}


//we return these as a pointer for a reason. Gender is mutable. Allow users to
//update these values *regardless* of if they identify as trans or not. 
var Man *genders.Gender = genders.NewGender(
  "man",
  genders.WithPublicGender(),
  genders.WithAllowedPronouns([]genders.Pronoun{HeHim}),
  )

var Woman *genders.Gender = genders.NewGender(
  "woman",
  genders.WithPublicGender(),
  genders.WithAllowedPronouns([]genders.Pronoun{SheHer}),
  )

var Nonbinary *genders.Gender = genders.NewGender(
  "non-binary",
  genders.WithPublicGender(),
  genders.WithAllowedPronouns([]genders.Pronoun{TheyThem}),
  )
