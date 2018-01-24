package main

import (
	"fmt"
	"os"

	dbmss "gitlab.10101111.com/oped/uparse.git/service"

	log "github.com/Sirupsen/logrus"
	"github.com/emirpasic/gods/lists/arraylist"
)

type LoadBalencer interface {
	InitData(map[string][]string, map[string]int)
	GetAgentByNewTopic(topic string) []string
	GetAgentByRmTopic(topic string) []string
}

type Facor struct {
	agent string
	fac   int
}

func FacCompare(a, b interface{}) int {
	fa := a.(*Facor)
	fb := b.(*Facor)
	switch {
	case fa.fac > fb.fac:
		return 1
	case fa.fac == fb.fac:
		return 0
	case fa.fac < fb.fac:
		return -1
	}
	return 0
}

type SerialLB struct {
	topicmap map[string][]string
	agentmap map[string]*Facor
	facor    *arraylist.List
}

func (lb *SerialLB) InitData(t map[string][]string, a map[string]int) {
	lb.topicmap = t
	lb.agentmap = make(map[string]*Facor)

	lb.facor = arraylist.New()
	for k, v := range a {
		faci := Facor{agent: k, fac: v}
		lb.facor.Add(&faci)
		lb.agentmap[k] = &faci
	}
	lb.facor.Sort(FacCompare)
}

func (lb *SerialLB) GetAgentByNewTopic(topic string) (ags []string) {

	if _, ok := lb.topicmap[topic]; ok {
		return
	}
	data, ok := lb.facor.Get(0)
	if !ok {
		return
	}
	agetn := data.(*Facor)
	ags = append(ags, agetn.agent)
	agetn.fac++
	lb.facor.Sort(FacCompare)

	return ags
}

func (lb *SerialLB) GetAgentByRmTopic(topic string) (ags []string) {

	if agslist, ok := lb.topicmap[topic]; !ok {
		return
	} else {
		ags = agslist
	}

	issorted := false
	for _, ag := range ags {
		if _, ok := lb.agentmap[ag]; ok {
			lb.agentmap[ag].fac--
			issorted = true
		}
	}
	if issorted {
		lb.facor.Sort(FacCompare)
	}
	return
}

type MoreLB struct {
	topicmap map[string][]string
	agentmap map[string]*Facor
	facor    *arraylist.List
	RetCount int
}

//  return RetCount agents whos load is smallest
func (lb *MoreLB) GetAgentByNewTopic(topic string) (ags []string) {
	if _, ok := lb.topicmap[topic]; ok {
		return
	}

	maxlen := lb.facor.Size()

	for i := 0; len(ags) < lb.RetCount && i < maxlen; i++ {
		data, ok := lb.facor.Get(i)
		if !ok {
			return
		}
		agetn := data.(*Facor)
		ags = append(ags, agetn.agent)
		agetn.fac++
	}
	lb.facor.Sort(FacCompare)
	return ags
}

func (lb *MoreLB) GetAgentByRmTopic(topic string) (ags []string) {
	if agslist, ok := lb.topicmap[topic]; !ok {
		return
	} else {
		ags = agslist
	}

	issorted := false
	for _, ag := range ags {
		if _, ok := lb.agentmap[ag]; ok {
			lb.agentmap[ag].fac--
			issorted = true
		}
	}
	if issorted {
		lb.facor.Sort(FacCompare)
	}
	return
}

func (lb *MoreLB) InitData(t map[string][]string, a map[string]int) {
	lb.topicmap = t
	lb.agentmap = make(map[string]*Facor)

	lb.facor = arraylist.New()
	for k, v := range a {
		faci := Facor{agent: k, fac: v}
		lb.facor.Add(&faci)
		lb.agentmap[k] = &faci
	}
	lb.facor.Sort(FacCompare)
}

func mydecoder(ag string, dm map[string]map[string]int, decoderlist []string) string {
	decoderscm, ok := dm[ag]
	if !ok {
		decoderscm = make(map[string]int)
		decoderscm[decoderlist[0]] = 1
		dm[ag] = decoderscm
		return decoderlist[0]
	}

	minindex := 0
	mintps := 9999
LOOP_XIAOMING:
	for pos, dec := range decoderlist {
		mytps, ok := decoderscm[dec]
		switch {
		case !ok: //this decoder is not used by uparse
			mintps = 0
			minindex = pos
			decoderscm[dec] = 0
			break LOOP_XIAOMING
		case mytps <= mintps: // find min topics who load
			mintps = mytps
			minindex = pos
		default:
			continue
		}

	}
	decoderscm[decoderlist[minindex]] += 1
	return decoderlist[minindex]
}

func initdecodersMap(los []*dbmss.TopicLogItem) map[string]map[string]int {
	decodersmap := make(map[string]map[string]int)
	for _, logi := range los {
		_, ok := decodersmap[logi.UparseAddr]
		if !ok {
			decodersmap[logi.UparseAddr] = make(map[string]int)
			decodersmap[logi.UparseAddr][logi.Decoder] = 1
		} else {
			_, ok = decodersmap[logi.UparseAddr][logi.Decoder]
			if !ok {
				decodersmap[logi.UparseAddr][logi.Decoder] = 1
			} else {
				decodersmap[logi.UparseAddr][logi.Decoder]++
			}
		}
	}
	return decodersmap
}

func main() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	data := map[string][]string{
		"test":  {"10.104,105.71:9999", "10.104.106.88:9999", "10.104.106.89:9999", "10.104.106.90:9999"},
		"zk":    {"10.104.106.88:9999", "10.104.106.89:9999", "10.104.106.90:9999"},
		"nginx": {"10.104.106.89:9999", "10.104.106.90:9999"},
	}
	data = make(map[string][]string)
	data["test"] = []string{"10.104,105.71:9999", "10.104.106.88:9999", "10.104.106.89:9999", "10.104.106.90:9999"}
	data["zk"] = []string{"10.104.106.89:9999", "10.104.106.90:9999"}
	data["nginx"] = []string{"10.104.106.89:9999", "10.104.106.90:9999"}

	adata := map[string]int{
		"10.104,105.71:9999": 1,
		"10.104.106.88:9999": 2,
		"10.104.106.89:9999": 3,
		"10.104.106.90:9999": 3,
	}

	//lb := &SerialLB{}
	lb := &MoreLB{RetCount: 5}
	lb.InitData(data, adata)
	log.Debugf("%s", lb.facor.String())
	arrays := lb.GetAgentByNewTopic("topicyyyyy")
	for _, ag := range arrays {
		log.Debugf("topicyyyyy consumed by %s", ag)
	}

	log.Debugf("%s", lb.facor.String())

	arrays = lb.GetAgentByRmTopic("test")
	for _, ag := range arrays {
		log.Debugf("test consumed by %s, removed it", ag)
	}
	log.Debugf("%s", lb.facor.String())

	arrays = lb.GetAgentByNewTopic("topicxxxxx")
	for _, ag := range arrays {
		log.Debugf("topicxxxxx consumed by %s", ag)
	}
	log.Debugf("%s", lb.facor.String())

	arrays = lb.GetAgentByNewTopic("topiczz")
	for _, ag := range arrays {
		log.Debugf("topiczz consumed by %s", ag)
	}

	log.Debugf("%s", lb.facor.String())
	arrays = lb.GetAgentByNewTopic("topicoooo")
	for _, ag := range arrays {
		log.Debugf("topicoooo consumed by %s", ag)
	}
	log.Debugf("%s", lb.facor.String())
	arrays = lb.GetAgentByNewTopic("topicaaaaaa")
	for _, ag := range arrays {
		log.Debugf("topicaaaaaaaa consumed by %s", ag)
	}

	log.Debugf("%s", lb.facor.String())
	arrays = lb.GetAgentByNewTopic("topicfffffff")
	for _, ag := range arrays {
		log.Debugf("topicffff consumed by %s", ag)
	}
	log.Debugf("%s", lb.facor.String())

	decodersmap := map[string]map[string]int{
		"host1": {"mutl1": 1, "mutl2": 1, "mutl3": 1},
		"host2": {"mutl1": 1, "mutl2": 3},
		"host3": {"mutl2": 4, "mutl3": 3},
	}
	decoders := []string{"mutl1", "mutl2", "mutl3"}

	var logss []*dbmss.TopicLogItem
	log1 := &dbmss.TopicLogItem{}
	log1.Tag = "ok"
	log1.UparseAddr = "host1"
	log1.Decoder = "mutl1"
	logss = append(logss, log1)

	log2 := &dbmss.TopicLogItem{}
	log2.Tag = "ok"
	log2.UparseAddr = "hosts2"
	log2.Decoder = "mutl2"
	logss = append(logss, log2)

	log3 := &dbmss.TopicLogItem{}
	log3.Tag = "ok"
	log3.UparseAddr = "hosts2"
	log3.Decoder = "mutl1"
	logss = append(logss, log3)

	log4 := &dbmss.TopicLogItem{}
	log4.Tag = "ok"
	log4.UparseAddr = "hosts2"
	log4.Decoder = "mutl1"
	logss = append(logss, log4)

	log5 := &dbmss.TopicLogItem{}
	log5.Tag = "ok"
	log5.UparseAddr = "hosts2"
	log5.Decoder = "mutl3"
	logss = append(logss, log5)

	decodersmap1 := initdecodersMap(logss)
	//decodersmap = decodersmap1
	fmt.Printf("map %v\n", decodersmap1)

	host := "host1"
	decc := mydecoder(host, decodersmap, decoders)
	fmt.Printf("%s decoder %s \n", host, decc)
	testok(decc, "mutl3")
	decc = mydecoder(host, decodersmap, decoders)
	fmt.Printf("%s decoder %s \n", host, decc)
	testok(decc, "mutl2")
	decc = mydecoder(host, decodersmap, decoders)
	fmt.Printf("%s decoder %s \n", host, decc)
	testok(decc, "mutl1")
	decc = mydecoder(host, decodersmap, decoders)
	fmt.Printf("%s decoder %s \n", host, decc)
	testok(decc, "mutl3")
	host = "host2"
	decc = mydecoder(host, decodersmap, decoders)
	fmt.Printf("%s decoder %s \n", host, decc)
	testok(decc, "mutl3")
	decc = mydecoder(host, decodersmap, decoders)
	fmt.Printf("%s decoder %s \n", host, decc)
	testok(decc, "mutl3")
	decc = mydecoder(host, decodersmap, decoders)
	fmt.Printf("%s decoder %s \n", host, decc)
	testok(decc, "mutl1")
	decc = mydecoder(host, decodersmap, decoders)
	fmt.Printf("%s decoder %s \n", host, decc)
	testok(decc, "mutl3")
	decc = mydecoder(host, decodersmap, decoders)
	fmt.Printf("%s decoder %s \n", host, decc)
	testok(decc, "mutl1")
	decc = mydecoder(host, decodersmap, decoders)
	fmt.Printf("%s decoder %s \n", host, decc)
	testok(decc, "mutl3")
	decc = mydecoder(host, decodersmap, decoders)
	fmt.Printf("%s decoder %s \n", host, decc)
	testok(decc, "mutl2")
	host = "host99"
	decc = mydecoder(host, decodersmap, decoders)
	fmt.Printf("%s decoder %s \n", host, decc)
	testok(decc, "mutl1")
	decc = mydecoder(host, decodersmap, decoders)
	fmt.Printf("%s decoder %s \n", host, decc)
	testok(decc, "mutl2")
	decc = mydecoder(host, decodersmap, decoders)
	fmt.Printf("%s decoder %s \n", host, decc)
	testok(decc, "mutl3")

}

func testok(zs, qw string) {
	if zs == qw {
		fmt.Printf("desired:%s result:%s, test ok  VVVVVVVVVVVVVVVVVVVVV\n", qw, zs)
	} else {
		fmt.Printf("desired:%s result:%s, test falure XXXXXXXXXXXXXXXXXXX \n", qw, zs)
	}
}
