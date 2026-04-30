package storefront

type Action struct {
	Label     string
	Href      string
	AriaLabel string
	Icon      string
}

type StorefrontPageProps struct {
	Hero        HeroProps
	Categories  []CategoryCard
	Collections []CollectionCard
	About       AboutProps
	Benefits    []BenefitItem
	Product     ProductPreview
	Rail        RailProps
	Footer      FooterProps
}

type HeroProps struct {
	Kicker       string
	Title        string
	Description  string
	Primary      Action
	SlideCurrent string
	SlideTotal   string
	Previous     Action
	Next         Action
	Image        ImageRef
}

type ImageRef struct {
	Src string
	Alt string
}

type CategoryCard struct {
	Title string
	Count string
	Href  string
	Image ImageRef
}

type CollectionCard struct {
	Label       string
	Title       string
	Description string
	Href        string
	Image       ImageRef
}

type AboutProps struct {
	Kicker      string
	Title       string
	Description string
	Items       []BenefitItem
}

type BenefitItem struct {
	Icon        string
	Title       string
	Description string
}

type ProductPreview struct {
	Badge          string
	Name           string
	Price          string
	OldPrice       string
	Description    string
	Gallery        []ImageRef
	MainImage      ImageRef
	ColorsLabel    string
	Colors         []SwatchItem
	SelectedColor  string
	Quantity       QuantityProps
	AddAction      Action
	FavoriteAction Action
	Details        []DetailRow
}

type SwatchItem struct {
	Label    string
	Value    string
	Class    string
	Selected bool
}

type QuantityProps struct {
	Label      string
	Value      string
	MinusLabel string
	PlusLabel  string
}

type DetailRow struct {
	Label string
	Value string
}

type RailProps struct {
	Cart        CartSummary
	Filter      FilterPanel
	Inspiration InspirationCard
	Newsletter  NewsletterProps
	Socials     []Action
}

type CartSummary struct {
	Title          string
	CountLabel     string
	Items          []CartItem
	TotalLabel     string
	Total          string
	CheckoutAction Action
	ViewCartAction Action
}

type CartItem struct {
	Title    string
	Subtitle string
	Price    string
	Image    ImageRef
}

type FilterPanel struct {
	Title         string
	ResetLabel    string
	CategoryLabel string
	CategoryValue string
	MaterialLabel string
	Materials     []SwatchItem
	Price         PriceRange
	ColorLabel    string
	Colors        []SwatchItem
	ShowAllAction Action
}

type PriceRange struct {
	Label    string
	MinLabel string
	MaxLabel string
	Value    string
}

type InspirationCard struct {
	Title       string
	Description string
	Action      Action
	Image       ImageRef
}

type NewsletterProps struct {
	Title       string
	Description string
	Placeholder string
	SubmitLabel string
}

type FooterProps struct {
	BrandName string
	Copyright string
	Groups    []FooterGroup
}

type FooterGroup struct {
	Title string
	Links []Action
}
