package iec61499

import (
	"errors"
	"strconv"
	"strings"
)

//GetArraySize returns the array size as an integer if there is one that can be parsed, otherwise 0
func (v Variable) GetArraySize() int {
	size, err := strconv.Atoi(v.ArraySize)
	if err != nil {
		return 0
	}
	return size
}

//GetInitialArray returns a formatted initial array if there is one to do so
func (v Variable) GetInitialArray() []string {
	//if cannot parse an array size then give up
	_, err := strconv.Atoi(v.ArraySize)
	if err != nil {
		return nil
	}

	//remove everything except commas and values
	raw := v.InitialValue
	raw = strings.TrimPrefix(raw, "[")
	raw = strings.TrimSuffix(raw, "]")

	raws := strings.Split(raw, ",")
	for i := 0; i < len(raws); i++ {
		raws[i] = strings.Trim(raws[i], " ")
	}
	return raws
}

//GetTransitionsForState returns only the transitions for a given source state
func (b *BasicFB) GetTransitionsForState(source string) []ECTransition {
	trans := make([]ECTransition, 0, len(b.Transistions))
	for i := 0; i < len(b.Transistions); i++ {
		if b.Transistions[i].Source == source {
			trans = append(trans, b.Transistions[i])
		}
	}
	return trans
}

//IsLoadFor used in templates for triggering updates of input data
func (e *Event) IsLoadFor(v *Variable) bool {
	for i := 0; i < len(e.With); i++ {
		if e.With[i].Var == v.Name {
			return true
		}
	}
	return false
}

//GetUniqueEventConnSources is used to list all unique sources for event connections
func (c *CompositeFB) GetUniqueEventConnSources() []string {
	sources := make([]string, 0, len(c.EventConnections)) //preallocate for speed
nextConn:
	for i := 0; i < len(c.EventConnections); i++ {
		//check if source already found
		for j := 0; j < len(sources); j++ {
			if sources[j] == c.EventConnections[i].Source {
				continue nextConn
			}
		}
		sources = append(sources, c.EventConnections[i].Source)
	}
	return sources
}

//GetUniqueDataConnSources is used to list all unique sources for data connections
func (c *CompositeFB) GetUniqueDataConnSources() []string {
	sources := make([]string, 0, len(c.DataConnections)) //preallocate for speed
nextConn:
	for i := 0; i < len(c.DataConnections); i++ {
		//check if source already found
		for j := 0; j < len(sources); j++ {
			if sources[j] == c.DataConnections[i].Source {
				continue nextConn
			}
		}
		sources = append(sources, c.DataConnections[i].Source)
	}
	return sources
}

//ConnectionWithType struct used in cases when we want to know what a Connection's Data Type is (this can store it)
type ConnectionWithType struct {
	Connection
	Type string
}

//SourceAndType is used in similiar cases to ConnectionWithType, when we want to know both the source's name and the data type (this can store it)
type SourceAndType struct {
	Source string
	Type   string
}

//GetUniqueDataConnSourcesWithTypes is used to list all unique sources for data connections and their types
func (f FB) GetUniqueDataConnSourcesWithTypes(otherBlocks []FB) ([]SourceAndType, error) {
	cats, err := f.GetDataConnectionTypes(otherBlocks)
	if err != nil {
		return nil, err
	}

	sources := make([]SourceAndType, 0, len(cats)) //preallocate for speed
nextConn:
	for i := 0; i < len(cats); i++ {
		//check if source already found
		for j := 0; j < len(sources); j++ {
			if sources[j].Source == cats[i].Source {
				continue nextConn
			}
		}
		sources = append(sources, SourceAndType{Source: cats[i].Source, Type: cats[i].Type})
	}
	return sources, nil
}

//GetDataConnectionTypes is used to list all data connections in a given FB with their types as well
func (f FB) GetDataConnectionTypes(otherBlocks []FB) ([]ConnectionWithType, error) {
	if f.CompositeFB == nil {
		return nil, nil //basic function blocks don't have dataconnections
	}

	c := f.CompositeFB
	connAndTypes := make([]ConnectionWithType, len(c.DataConnections))

conns:
	for i := 0; i < len(c.DataConnections); i++ {
		//store all connection data
		connAndTypes[i].Connection = c.DataConnections[i]

		conn := &connAndTypes[i].Connection

		found := false

		//find the type based only off the source
		if !strings.Contains(conn.Source, ".") {
			//source is from this block's parent port
			if f.InputVars != nil {
				for j := 0; j < len(f.InputVars.Variables); j++ {
					if f.InputVars.Variables[j].Name == conn.Source {
						found = true
						connAndTypes[i].Type = f.InputVars.Variables[j].Type
						continue conns
					}
				}
			}
		}

		//still here? source must be from a child block's output port
		splitSourceName := strings.Split(conn.Source, ".")
		if len(splitSourceName) != 2 {
			return nil, errors.New("Source of dataconnection '" + conn.Source + "' has an incorrect number of periods (should be 0 or 1).")
		}
		childName := splitSourceName[0]
		sourceName := splitSourceName[1]
		childType := ""

		//find the child's real block name
		for j := 0; j < len(c.FBs); j++ {
			if c.FBs[j].Name == childName {
				childType = c.FBs[j].Type
			}
		}
		if childType == "" {
			return nil, errors.New("Could not find source of dataconnection '" + conn.Source + "' as child block can't be found.")
		}

		//scan through all blocks trying to find correct API type
		for j := 0; j < len(otherBlocks); j++ {
			if otherBlocks[j].Name == childType { //matched, now scan their API
				if otherBlocks[j].OutputVars != nil {
					for k := 0; k < len(otherBlocks[j].OutputVars.Variables); k++ {
						if otherBlocks[j].OutputVars.Variables[k].Name == sourceName {
							found = true
							connAndTypes[i].Type = otherBlocks[j].OutputVars.Variables[k].Type
							continue conns
						}
					}
				} else {
					return nil, errors.New("Source of dataconnection '" + conn.Source + "' has no output vars!")
				}
			}
		}

		if found == false {
			return nil, errors.New("Could not find source of dataconnection '" + conn.Source + "' in any included file.")
		}
	}

	return connAndTypes, nil
}
