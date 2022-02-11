package component

// ComponentInfoImpl holds information of the component and it's dependencies when seting up the platform
type ComponentInfoImpl struct {
	Name                    string
	Description             string
	Group                   string //enum
	DependsOn               []ComponentInfoImpl
	DependsOnGroup          string //enum
	RequiresInformationFrom []ComponentInfoImpl
}

const (
	ManagementBootstrap = "Management Bootstrap"
)

// DefaultComponentInfoImpl should be used for initializing the ComponentInfoImpl so whenever we add
// new things to the object not all the code breaks at once, we can gradually add the desired information
func DefaultComponentInfoImpl() ComponentInfoImpl {
	return ComponentInfoImpl{
		Name:                    "Change me",
		Description:             "",
		Group:                   "Change me",
		DependsOn:               []ComponentInfoImpl{},
		DependsOnGroup:          "",
		RequiresInformationFrom: []ComponentInfoImpl{},
	}
}

// ComponentInfo is an interface that allows for access to ComponentInfoImpl
type ComponentInfo interface {
	Info() ComponentInfoImpl
}

// ------------------------ Remove below
// ComponentDepInfo holds the description of the component and it's dependencies when seting up the platform
type ComponentDepInfo struct {
	Name           string
	Description    string
	Group          string
	DependsOn      []string
	DependsOnGroup string
}

// ComponentInfo is an interface that allows for access to ComponentDepInfo.
type ComponentInfoTwo interface {
	Info() ComponentDepInfo
}
