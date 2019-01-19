package main

import (
	"bufio"
	"fmt"
	"github.com/ghodss/yaml"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"os"
)

//
//import (
//	"fmt"
//	"io/ioutil"
//	"k8s.io/apimachinery/pkg/api/resource"
//	"k8s.io/apimachinery/pkg/runtime"
//	"k8s.io/client-go/kubernetes/scheme"
//	"k8s.io/klog"
//	"log"
//	"regexp"
//	"strings"
//	"k8s.io/cli-runtime/pkg/genericclioptions/resource/builder"
//	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
//	"k8s.io/cli-runtime/pkg/genericclioptions"
//
//
//)
//
//func parseK8sYaml(fileR []byte) []runtime.Object {
//
//	acceptedK8sTypes := regexp.MustCompile(`(Role|ClusterRole|RoleBinding|ClusterRoleBinding|ServiceAccount)`)
//	fileAsString := string(fileR[:])
//	sepYamlfiles := strings.Split(fileAsString, "---")
//	retVal := make([]runtime.Object, 0, len(sepYamlfiles))
//	for _, f := range sepYamlfiles {
//		if f == "\n" || f == "" {
//			// ignore empty cases
//			continue
//		}
//
//		decode := scheme.Codecs.UniversalDeserializer().Decode
//		obj, groupVersionKind, err := decode([]byte(f), nil, nil)
//
//		if err != nil {
//			log.Println(fmt.Sprintf("Error while decoding YAML object. Err was: %s", err))
//			continue
//		}
//
//		if !acceptedK8sTypes.MatchString(groupVersionKind.Kind) {
//			log.Printf("The custom-roles configMap contained K8s object types which are not supported! Skipping object with type: %s", groupVersionKind.Kind)
//		} else {
//			retVal = append(retVal, obj)
//		}
//
//	}
//	return retVal
//}
//
//func check(e error) {
//	if e != nil {
//		panic(e)
//	}
//}
//
//func likeKubectl() {
//
//		cmdNamespace := "default"
//		enforceNamespace := false
//
//	kubeConfigFlags := genericclioptions.NewConfigFlags()
//
//	matchVersionKubeConfigFlags := cmdutil.NewMatchVersionFlags(kubeConfigFlags)
//	f := cmdutil.NewFactory(matchVersionKubeConfigFlags)
//
//		r := f.NewBuilder().
//		Unstructured().
//		Schema(schema).
//		ContinueOnError().
//		NamespaceParam(cmdNamespace).DefaultNamespace().
//		FilenameParam(enforceNamespace, &o.FilenameOptions).
//		LabelSelectorParam(o.Selector).
//		Flatten().
//		Do()
//		err = r.Err()
//		if err != nil {
//		return err
//	}
//
//		count := 0
//		err = r.Visit(func(info *resource.Info, err error) error {
//		if err != nil {
//		return err
//	}
//		if err := kubectl.CreateOrUpdateAnnotation(cmdutil.GetFlagBool(cmd, cmdutil.ApplyAnnotationsFlag), info.Object, scheme.DefaultJSONEncoder()); err != nil {
//		return cmdutil.AddSourceToErr("creating", info.Source, err)
//	}
//
//		if err := o.Recorder.Record(info.Object); err != nil {
//		klog.V(4).Infof("error recording current command: %v", err)
//	}
//
//		if !o.DryRun {
//		if err := createAndRefresh(info); err != nil {
//		return cmdutil.AddSourceToErr("creating", info.Source, err)
//	}
//	}
//
//		count++
//
//		return o.PrintObj(info.Object)
//	})
//		if err != nil {
//		return err
//	}
//		if count == 0 {
//		return fmt.Errorf("no objects passed to create")
//	}
//		return nil
//	}
//}
//
func main() {

	//yams, err := ioutil.ReadFile("./yaml/istio-crds-v0.3.0.yaml")

	//j, err := yaml.YAMLToJSON([]byte(yams))
	//if err != nil {
	//	fmt.Print(err)
	//}
	//fmt.Printf("json: %s", string(j))

	//unstructured.UnstructuredJSONScheme.Decode()

	//fmt.Printf("Unstructured: %v, %v", u, u.IsList())

	//	yaml.

	f, err := os.Open("./yaml/istio-crds-v0.3.0.yaml")
	if err != nil {
		fmt.Print(err)
	}

	var obj []byte

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())

		l := scanner.Text()
		if l == "---" {
			u := unstructured.Unstructured{}
			err = yaml.Unmarshal(obj, &u)
			if err != nil {
				fmt.Print(err)
			}
			fmt.Printf("Unstructured: %v\n", u)

			// clean
			obj = []byte{}
		} else {
			obj = append(obj, scanner.Bytes()...)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Print(err)
	}

	//b1 := make([]byte, 5)
	//n1, err := f.Read(b1)
	//if err != nil {
	//	fmt.Print(err)
	//}
	//fmt.Printf("%d bytes: %s\n", n1, string(b1))

}
