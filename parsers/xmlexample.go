package parsers

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type TopLevelServers struct {
	XMLName        xml.Name `xml:"servers"`
	Version        string   `xml:"version,attr"`
	ChildrenServer []Server `xml:"server"`
	Description    string   `xml:",innerxml"`
}

type Server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

func Parse() {
	v := TopLevelServers{}
	data, err := ReadFile("parsers/servers.xml")
	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}
	fmt.Println("Name:", v.XMLName, "Version: ", v.Version, "Desc: ", v.Description, "childrens: ", v.ChildrenServer)
}

func GenXML() {
	v := &TopLevelServers{Version: "1", Description: "\n\tThis is an xml description for servers list."}
	v.ChildrenServer = append(v.ChildrenServer, Server{ServerName: "Singapore_VPN", ServerIP: "127.0.0.1"})
	v.ChildrenServer = append(v.ChildrenServer, Server{ServerName: "Banglore_VPN", ServerIP: "127.0.0.2"})
	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}

func ReadFile(filename string) ([]byte, error) {
	file, err := os.Open(filename) // for read access
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil, err
	}
	return data, nil
}
