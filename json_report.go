package main
import (
	"encoding/json"
	"slices"
	"os"
	"fmt"
)

func writeJSONReport(pages map[string]PageData, filename string) error{
	keys:=make([]string,0,len(pages))
	for key:= range pages{
		keys=append(keys,key)
	}
	slices.Sort(keys)
	pageDataSlice:=make([]PageData,len(keys))
	for i,key:= range keys{
		pageDataSlice[i]=pages[key]
	}
	data, err := json.MarshalIndent(pageDataSlice, "", "  ")
	if err != nil{
		return fmt.Errorf("Marshal pageData Slice err:%w\n",err)
	}
	if err:= os.WriteFile(filename, data, 0644);err != nil{
		return fmt.Errorf("Write file :%s, err:%w\n",filename,err)
	}
	return nil
}