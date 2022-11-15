package genders

//This is just an example library implementing a trans and gender-divergent friendly
//data structure for handling information about an individual, in the rare case where
//it is necessary. Here's an important question to ask: do you actually need to know
//this data? What happens if you don't know? Does it matter?
type Gender struct {
  //Pronouns that express the individual's gender experience
  Pronouns Pronouns
  //Is this information public or private?
  IsPublic bool
  //Name is the name of the gender, if the individual wishes to assosciate with one
  Name string
  //Assosciates is a list of other genders that the individual assosciates their gender with.
  //For instance, someone may be nonbinary and also assosciate with manhood.
  Associates map[string]Gender
}

//Instantiates a new Gender 
func NewGender(name string, opts...GenderOption) *Gender {
  gndr := &Gender{
    Pronouns: Pronouns{},
    IsPublic: false,
    Name: "",
    Associates: map[string]Gender{},
  }
  gndr.Name = name

  for _,opt := range opts {
    opt(gndr)
  }

  return gndr
}

//Pronouns is a meta-struct which contains lists of pronouns categorized by their
//relative usage by a person. Primary are interchangably usable, allowed are okay but
//not necessarily preferred, and disallowed are never to be used. These are implmented 
//as a list of Pronoun structs as an individual can use multiple.
type Pronouns struct {
  //List of primary pronouns, meaning those which are welcomed and interchangably usable
  //when referring to a person
  Primary []Pronoun
  //List of allowed but not necessarily preferred pronouns for a person
  Allowed []Pronoun
  //List of pronouns that are not welcomed by a given individual
  Disallowed []Pronoun
  //Flag which signifies that an individual is okay with the use of any pronouns.
  //If this is used AND disallowed is also populated, the interpretation should be
  //"all except"
  AllowAny bool
}

//Pronoun is a basic gramatical data structure describing the different potential
//words that can be used to refer to an individual given the kind of pronoun. We
//include things like first person, second plural, etc. as not all languages have
//neutral first-person or plural pronouns (eg, in Japanese there are gendered first-person
//pronouns [watashi 私、boku 僕、ore 俺]). I am not a native speaker of anything but english,
//so I will leave things like declensions and other more complex gramatical rules to someone
//with experience in a language which uses them. We also include fourth and fifth persons, as
//some languages such as Ojibwe implement them.
type Pronoun struct {
  //The language in which this pronoun is used
  Language string
  //A short-hand reference for the set of pronouns that are gramatically similar/the same
  //that would reside together in this data structure, eg 'they/them'
  Alias string
  //The first person pronoun for the given grouping
  FirstPerson PronounPart
  //Plural first person pronoun for the given grouping
  FirstPersonPlural PronounPart
  //Second person pronoun for the grouping
  SecondPerson PronounPart
  //Plural second person pronoun for the grouping
  SecondPersonPlural PronounPart
  //Third person pronoun for the grouping
  ThirdPerson PronounPart
  //Plural third person pronoun for the grouping
  ThirdPersonPlural PronounPart
  //Fourth person pronoun for the grouping
  FourthPerson PronounPart
  //Fourth person plural pronoun for the grouping
  FourthPersonPlural PronounPart
  //Fifth person pronoun for the grouping
  FifthPerson PronounPart
  //Fifth person plural pronoun for the grouping
  FifthPersonPlural PronounPart
}

//PronounPart holds the actual string value of a pronoun given its gramatical use-case.
type PronounPart struct {
  Nominative string
  Genitive string
  Dative string
  Accusative string
  Ablative string
  Prepositional string
  Instrumental string
  GenitiveDeterminer string
  Reflexive string
}

//GenderOption provides generator handlers for NewGender()
type GenderOption func(*Gender)

//Sets a list of primary pronouns used by the particular person
func WithPrimaryPronouns(p []Pronoun) GenderOption {
  return func(g *Gender) {
    g.Pronouns.Primary = p
  }
}

func WithAllowedPronouns(p []Pronoun) GenderOption {
  return func(g *Gender) {
    g.Pronouns.Allowed = p
  }
}

func WithDisallowedPronouns(p []Pronoun) GenderOption {
  return func(g *Gender) {
    g.Pronouns.Disallowed = p
  }
}

func WithAnyPronouns() GenderOption {
  return func(g *Gender) {
    g.Pronouns.AllowAny = true
  }
}

func WithPublicGender() GenderOption {
  return func(g *Gender) {
    g.IsPublic = true
  }
}

func WithName(name string) GenderOption {
  return func(g *Gender) {
    g.Name = name
  }
}

func WithAssociates(asocs map[string]Gender) GenderOption {
  return func(g *Gender) {
    g.Associates = asocs
  }
}


