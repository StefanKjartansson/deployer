package deployer

var (
	Projects = make(map[string]Project)
)

type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
    Provisioner string `json:"provisioner"`
}


func NewProject(name string) (p Project) {
	p.ID = Uuid()
	p.Name = name
    p.Provisioner = `
git status
uptime
    `
	Projects[p.ID] = p
	return
}

