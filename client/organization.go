package client

import (
	"log"

	"github.com/lxn/walk"
)

type Organization struct {
	orgId    string
	orgName  string
	parent   *Organization
	children []*Organization
}

func NewOrganization(id string, orgName string, parent *Organization) *Organization {
	return &Organization{orgId: id, orgName: orgName, parent: parent}
}

var _ walk.TreeItem = new(Organization)

func (org *Organization) Text() string {
	return org.orgName
}

func (org *Organization) Parent() walk.TreeItem {
	if org.parent == nil {
		return nil
	}
	return org.parent
}

func (org *Organization) ChildCount() int {
	if org.children == nil {
		if err := org.ResetChildren(); err != nil {
			log.Print(err)
		}
	}
	return len(org.children)
}

func (org *Organization) ChildAt(index int) walk.TreeItem {
	return org.children[index]
}

func (org *Organization) ResetChildren() error {
	/*org.children = nil

	orgs, err := (&dataSource.OrganizationEntity{Id: org.orgId}).ListSubOrg()
	if err != nil {
		return err
	}

	if len(orgs) == 0 {
		org.children = []*Organization{}
		return nil
	}

	for i := 0; i < len(orgs); i++ {
		if orgs[i] == nil {
			continue
		}

		org.children = append(org.children, NewOrganization(orgs[i].Id, orgs[i].OrgName, org))
	}*/

	return nil
}

type OrganizationTreeModel struct {
	walk.TreeModelBase
	roots []*Organization
}

var _ walk.TreeModel = new(OrganizationTreeModel)

func (*OrganizationTreeModel) LazyPopulation() bool {
	return false
}

func (m *OrganizationTreeModel) RootCount() int {
	return len(m.roots)
}

func (m *OrganizationTreeModel) RootAt(index int) walk.TreeItem {
	return m.roots[index]
}
