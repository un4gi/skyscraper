package models

const (
	ANALYTICS       string = "listing_analytics/categories?listing_id="
	STATUS_ACCEPTED string = "&status=accepted"
	STATUS_IN_QUEUE string = "&status=in_queue"
	STATUS_REJECTED string = "&status=rejected"
	SYNACKAPI       string = "https://platform.synack.com/api/"
	TARGETS 	string = "https://platform.synack.com/api/targets?filter%5Bprimary%5D=registered&filter%5Bindustry%5D=all&sorting%5Bfield%5D=dateUpdated&sorting%5Bdirection%5D=desc"
)

var (
	Url string = SYNACKAPI + ANALYTICS + "id-placeholder" + STATUS_ACCEPTED
)

type Analytics struct {
	ListingID string `json:"listing_id"`
	Type      string `json:"type"`
	Value     []struct {
		Categories           []string `json:"categories"`
		ExploitableLocations []struct {
			Type      string `json:"type"`
			Value     string `json:"value"`
			CreatedAt int    `json:"created_at"`
			Status    string `json:"status"`
		} `json:"exploitable_locations"`
	} `json:"value"`
}

type TargetData []struct {
	Category struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"category"`
	Organization struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"organization"`
	Codename  string `json:"codename"`
	Slug      string `json:"slug"`
	OutageWin []struct {
		StartDate   int `json:"start_date"`
		EndDate     int `json:"end_date"`
		OutageStart int `json:"outage_starts_on"`
		OutageEnds  int `json:"outage_ends_on"`
		Options     struct {
			Days      []int  `json:"days"`
			Frequency string `json:"frequency"`
		} `json:"options"`
		WindowActive bool `json:"is_window_active"`
	} `json:"outage_windows"`
	SRT_Notes   string   `json:"srt_notes"`
	DateUpdated int      `json:"dateUpdated"`
	Active      bool     `json:"isActive"`
	New         bool     `json:"isNew"`
	Registered  bool     `json:"isRegistered"`
	Name        string   `json:"name"`
	AvgPayout   float64  `json:"averagePayout"`
	LastSubm    int      `json:"lastSubmitted"`
	StartDate   int      `json:"start_date"`
	EndDate     int      `json:"end_date"`
	VulnDisc    bool     `json:"vulnerability_discovery"`
	Workspace   bool     `json:"workspace_access_missing"`
	Updated     bool     `json:"isUpdated"`
	Incentives  []string `json:"incentives"`
}
