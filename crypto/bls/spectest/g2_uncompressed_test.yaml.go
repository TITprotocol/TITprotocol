// Code generated by yaml_to_go. DO NOT EDIT.
// source: g2_uncompressed.yaml

package spectest

type G2UncompressedTest struct {
	Title         string   `json:"title"`
	Summary       string   `json:"summary"`
	ForksTimeline string   `json:"forks_timeline"`
	Forks         []string `json:"forks"`
	Config        string   `json:"config"`
	Runner        string   `json:"runner"`
	Handler       string   `json:"handler"`
	TestCases     []struct {
		Input struct {
			Message []byte `json:"message" ssz:"size=32"`
			Domain  uint64 `json:"domain"`
		} `json:"input"`
		Output [][][]byte `json:"output"`
	} `json:"test_cases"`
}
