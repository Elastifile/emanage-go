package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	. "github.com/onsi/gomega"

	"helputils"
	"logging"
	"types"
)

// func TestFullTomlFile(t *testing.T) {
// 	confStr := `
// 	[setup]
// 	loaders = [	"loaderA", "loaderB", "loaderC", "loaderD" ]
// 	vheads = [ "vheadA", "vheadB", "vheadC", "vheadD" ]

// 	[docker]
// 	local = true
// 	port = 2375
// 	registry = "registry.il.elastifile.com"

// 	[docker.push_pull]
// 	use_personal = true

// 	[elfs]
// 	frontend = "192.168.0.1"
// 	export = "my_fs0/root_0"

// 	[emanage]
// 	server = "10.11.16.187"

// 	[ecs]
// 	server = "10.11.16.187:10016"

// 	[vcenter]
// 	host = "vc8c.lab.il.elastifile.com"
// 	username = "root"
// 	password = "MyPassword"

// 	[system.deploy.data.datanet]
// 	address = "8.0.0.0"
// 	subnet = 24
// 	name = "Elastifile-DATA"

// 	[system.deploy.data.nfsnet]
// 	address = "192.168.0.1"
// 	subnet = 24
// 	name = "Elastifile-CLN"

// 	[system.deploy]
// 	production_mode = true
// 	configfile = "infra/system_deploy/setups_config/vc-8c/qa-app-vc-8c-3Nodes.json"
// 	forceinstall = true
// 	exitifpanic = true
// 	replication_level = 2

// 	[system.deploy.data]
// 	externalNetName = "VM Network"
// 	disksPrefix = ["RealSSD"]
// 	vlanid = 8
// 	enodes = [
// 		"vHead-esx8a.lab.il.elastifile.com",
// 		"vHead-esx8b.lab.il.elastifile.com",
// 		"vHead-esx8c.lab.il.elastifile.com",
// 		"vHead-esx8d.lab.il.elastifile.com",
// 	]
// 	`

// 	config := confFromStr(t, confStr)
// 	dump(t, config)
// }

// func TestPartialTomlFile(t *testing.T) {
// 	confStr := `
// 	[system.deploy.data]
// 	vlanid = 25
// 	vlanid2 = 1025

// 	[system.deploy.data.datanet]
// 	address = "5.0.0.1"
// 	subnet = 24
// 	name = "Elastifile-DATA"

// 	[system.deploy.data.datanet2]
// 	address = "5.10.0.0"
// 	subnet = 24
// 	name = "Elastifile-DATA-2"
// 	`

// 	config := confFromStr(t, confStr)
// 	dump(t, config.System.Deploy.Data)
// }

func TestParseImages(t *testing.T) {
	DeployConfig = `
{
    "agent/sanity": "registry.il.elastifile.com/agent/sanity:265d9f95b8cf9945b6e64b04f8c56ba5208da8a7", 
    "agent/slave": "registry.il.elastifile.com/agent/slave:db162802721046c29dda90a41b9213589e535059", 
    "agent/tooly": "registry.il.elastifile.com/agent/tooly:6e64b1c937b23ede6a36e19fefedf67773c86111", 
    "system/ecs": "registry.il.elastifile.com/system/ecs:4ce0c02420e8a644f73f8fc380bcf41aace66088", 
    "system/elfs": "registry.il.elastifile.com/system/elfs:8efd933e2e88c974a9c1d9f59e5519bfac4f80e8", 
    "system/emanage": "registry.il.elastifile.com/system/emanage:396d2a4ef575188857d33adf75809149603eba13", 
    "system/emanagedb": "registry.il.elastifile.com/system/emanagedb:cc6e916342cd6a897c4db6d843f2f00795c5672e", 
    "test/newapi": "registry.il.elastifile.com/test/newapi:5a5a18fe58e6f5ee13afd46dfe146b95b824c818", 
    "thirdparty/gnatsd": "registry.il.elastifile.com/thirdparty/gnatsd:7083a5b7d23ff0ceda1205fca21197fb8a29acd4", 
    "thirdparty/minio": "registry.il.elastifile.com/thirdparty/minio:b2bccb2c610e86cc5091807af816f3e6ec6c6836", 
    "tool/cthon": "registry.il.elastifile.com/tool/cthon:f50f6b6922dd21d610bf234c41bf7b8a7c1c7d48", 
    "tool/erun": "registry.il.elastifile.com/tool/erun:4d61d911b7a8a5ebe6f021ad810c64df8e41dd5c", 
    "tool/fstool": "registry.il.elastifile.com/tool/fstool:6dd244a2e3090dbac94f7a1ff633c45d8b879702", 
    "tool/fio": "registry.il.elastifile.com/tool/fio:346926fd82cdd41f10b7e2efea93423b22cfccc2", 
    "tool/sfs2008": "registry.il.elastifile.com/tool/sfs2008:9ecd45c14aaaea58e47e918b5ac437bb2642419d", 
    "tool/sfs2014": "registry.il.elastifile.com/tool/sfs2014:5458df4277691715be9a660001ebc3f9d11532f9", 
    "tool/vdbench": "registry.il.elastifile.com/tool/vdbench:df90bfc784228b53dd17a0272dac991d212e6499"
}
`
	actual := GetImage(VDBenchImage).String()
	expected := "registry.il.elastifile.com/tool/vdbench:df90bfc784228b53dd17a0272dac991d212e6499"
	if actual != expected {
		t.Fatalf("Incorrectly parsed image name. Expected %s, got %s", expected, actual)
	}
}

func TestConfFromEnv(t *testing.T) {
	conf, err := ConfigFromEnvironment()
	if err != nil {
		t.Fatal(err)
	}
	spew.Dump(conf)
}

// confFromStr creates temporary toml file from given string and returns config loaded from that file.
func confFromStr(t *testing.T, confStr string) types.Config {
	f, err := toFile(t, confStr)
	check(t, err)
	defer func() { _ = os.Remove(f.Name()) }()

	conf := types.NewConfig()

	options := []string{}
	err = FromAllSources(nil, nil, options)
	check(t, err)
	return *conf
}

func toFile(t *testing.T, str string) (*os.File, error) {
	f, err := ioutil.TempFile("", "configTest")
	if err != nil {
		return nil, err
	}

	err = ioutil.WriteFile(f.Name(), []byte(str), os.ModePerm)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// pretty print data structures (using pkg 'spew')
func dump(t *testing.T, any interface{}) {
	t.Log(spew.Sdump(any))
}

func check(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func TestShouldWorkaround(t *testing.T) {
	var logger = logging.NewLogger("config")
	logging.Setup(nil)
	if ShouldWorkaround(nil, "ISSUE-1234", logger, "doing it anyways", "this", "that") {
		t.Log("DOING IT!")
	}
}

func TestEnsuredHeader(t *testing.T) {
	RegisterTestingT(t)

	tmpJson, err := helputils.TmpRandomFileName("ensured-header")
	check(t, err)
	tmpJson += ".json"
	check(t, ioutil.WriteFile(tmpJson, []byte(`
{
	"data": {
		"emanage": []
	}
}`), os.ModePerm))
	// defer func() { os.Remove(tmpJson) }()

	ensuredJson, err := ensureSystemElabParents(tmpJson)
	Expect(err).NotTo(HaveOccurred())
	// defer func() { os.Remove(ensuredJson) }()

	Expect(ensuredJson).NotTo(Equal(tmpJson))
	t.Log("json:", tmpJson, "ensured:", ensuredJson)

	body, err := helputils.ReadAll(ensuredJson)
	Expect(err).NotTo(HaveOccurred())

	var f interface{}
	err = json.Unmarshal(body, &f)
	body, err = json.Marshal(f)
	Expect(err).NotTo(HaveOccurred())
	t.Logf("%s", body)
}
