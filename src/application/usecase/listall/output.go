package listall

type Output interface {
	output(data []string) []string
}
