package xlripper

import (
	"encoding/json"
	"testing"

	"github.com/bitflip-software/xlripper/xmlprivate"
)

type strPair struct {
	a string
	b string
}

func TestUnitCellXML(t *testing.T) {
	for ix, pair := range pairs {
		got, err := xmlprivate.ParseCellXML(pair.a)

		if err != nil {
			t.Errorf("test index %d: error parsing the 'want' value from json: %s", ix, err.Error())
			continue
		}

		want := xmlprivate.CellXML{}
		err = json.Unmarshal([]byte(pair.b), &want)

		if err != nil {
			t.Errorf("test index %d: error parsing the 'want' value from json: %s", ix, err.Error())
			continue
		}

		gotJS, _ := json.Marshal(got)
		wantJS, _ := json.Marshal(want)
		if got != want {
			t.Error(tfail(t.Name(), "got, err := xmlprivate.ParseCellXML(pair.a)", string(gotJS), string(wantJS)))
		}
	}
}

func TestUnitCellCoreXML(t *testing.T) {
	for ix, pair := range pairs {
		want, err := xmlprivate.ParseCellXML(pair.a)

		if err != nil {
			t.Errorf("test index %d: error parsing the 'want' value from json: %s", ix, err.Error())
			continue
		}

		ccxFromJSON := cellCoreXML{}
		err = ccxFromJSON.UnmarshalJSON([]byte(pair.b))

		if err != nil {
			t.Errorf("test index %d: error during cellCoreXML.UnmarshalJSON: %s", ix, err.Error())
			continue
		}

		ccxFromXML := cellCoreXML{}
		err = ccxFromXML.parseXML([]rune(pair.a))

		if err != nil {
			t.Errorf("test index %d: error during cellCoreXML.parseXML: %s", ix, err.Error())
			continue
		}

		got := ccxFromJSON.x
		gotJS, _ := json.Marshal(got)
		wantJS, _ := json.Marshal(want)

		if want != want {
			t.Error(tfail(t.Name(), "want, err := xmlprivate.ParseCellXML(pair.a)", string(gotJS), string(wantJS)))
		}

		got = ccxFromXML.x
		gotJS, _ = json.Marshal(got)
		wantJS, _ = json.Marshal(want)

		if want != want {
			t.Error(tfail(t.Name(), "want, err := xmlprivate.ParseCellXML(pair.a)", string(gotJS), string(wantJS)))
		}

		stmt := "ccxFromXML.value()"
		gotS := ccxFromXML.value()
		wantS := ""

		if want.T == "inlineStr" {
			wantS = want.InlineString.Str
		} else {
			wantS = want.V
		}

		if *gotS != wantS {
			t.Error(tfail(t.Name(), stmt, *gotS, wantS))
		}

		stmt = "ccxFromXML.valueRunes()"
		gotR := ccxFromXML.valueRunes()
		wantR := []rune("")

		if want.T == "inlineStr" {
			wantS = want.InlineString.Str
		} else {
			wantS = want.V
		}

		if *gotS != wantS {
			t.Error(tfail(t.Name(), stmt, string(gotR), string(wantR)))
		}

		stmt = "ccxFromXML.cellReference()"
		*gotS = ccxFromXML.cellReference()
		wantS = want.R

		if *gotS != wantS {
			t.Error(tfail(t.Name(), stmt, *gotS, wantS))
		}

		stmt = "ccxFromXML.cellReferenceRunes()"
		gotR = ccxFromXML.cellReferenceRunes()
		wantR = []rune("")

		if *gotS != wantS {
			t.Error(tfail(t.Name(), stmt, string(gotR), string(wantR)))
		}

		stmt = "ccxFromXML.typeInfo()"
		*gotS = ccxFromXML.typeInfo().String()
		wantS = want.T

		if *gotS != wantS {
			t.Error(tfail(t.Name(), stmt, *gotS, wantS))
		}
	}
}

var pairs = []strPair{
	{`<x:c r="AC182" s="0"><x:v>0.00082725</x:v></x:c>`, `{"ref":"AC182","type":"","value":"0.00082725","inline_string":{"string_value":""}}`},
	{`<x:c r="P10271" s="0" t="s"><x:v>69</x:v></x:c>`, `{"ref":"P10271","type":"s","value":"69","inline_string":{"string_value":""}}`},
	{`<x:c r="AD12116" s="0"><x:v>0.00046727</x:v></x:c>`, `{"ref":"AD12116","type":"","value":"0.00046727","inline_string":{"string_value":""}}`},
	{`<x:c r="X15616" s="0"><x:v>0.0010</x:v></x:c>`, `{"ref":"X15616","type":"","value":"0.0010","inline_string":{"string_value":""}}`},
	{`<x:c r="BC17916" s="0"><x:v>0.00007357</x:v></x:c>`, `{"ref":"BC17916","type":"","value":"0.00007357","inline_string":{"string_value":""}}`},
	{`<x:c r="AJ21215" s="0"><x:v>0.001530</x:v></x:c>`, `{"ref":"AJ21215","type":"","value":"0.001530","inline_string":{"string_value":""}}`},
	{`<x:c r="BH21002" s="0"><x:v>0.000350</x:v></x:c>`, `{"ref":"BH21002","type":"","value":"0.000350","inline_string":{"string_value":""}}`},
	{`<x:c r="E17821" s="0" t="s"><x:v>768</x:v></x:c>`, `{"ref":"E17821","type":"s","value":"768","inline_string":{"string_value":""}}`},
	{`<x:c r="I25482" s="0" t="s"><x:v>449</x:v></x:c>`, `{"ref":"I25482","type":"s","value":"449","inline_string":{"string_value":""}}`},
	{`<x:c r="G19921" s="3"><x:v>6282</x:v></x:c>`, `{"ref":"G19921","type":"","value":"6282","inline_string":{"string_value":""}}`},
	{`<x:c r="AY29601" s="0" t="s"><x:v>85</x:v></x:c>`, `{"ref":"AY29601","type":"s","value":"85","inline_string":{"string_value":""}}`},
	{`<x:c r="AF29633" s="0" t="s"><x:v>71</x:v></x:c>`, `{"ref":"AF29633","type":"s","value":"71","inline_string":{"string_value":""}}`},
	{`<x:c r="AO29493" s="0"><x:v>0</x:v></x:c>`, `{"ref":"AO29493","type":"","value":"0","inline_string":{"string_value":""}}`},
	{`<x:c r="AC21606" s="0"><x:v>0.00007272</x:v></x:c>`, `{"ref":"AC21606","type":"","value":"0.00007272","inline_string":{"string_value":""}}`},
	{`<x:c r="AZ33435" s="0" t="s"><x:v>73</x:v></x:c>`, `{"ref":"AZ33435","type":"s","value":"73","inline_string":{"string_value":""}}`},
	{`<x:c r="AP34386" s="0"><x:v>0.9999</x:v></x:c>`, `{"ref":"AP34386","type":"","value":"0.9999","inline_string":{"string_value":""}}`},
	{`<x:c r="BA38152" s="0" t="s"><x:v>74</x:v></x:c>`, `{"ref":"BA38152","type":"s","value":"74","inline_string":{"string_value":""}}`},
	{`<x:c r="AC39540" s="0"><x:v>0.00081889</x:v></x:c>`, `{"ref":"AC39540","type":"","value":"0.00081889","inline_string":{"string_value":""}}`},
	{`<x:c r="Y39494" s="0" t="s"><x:v>100</x:v></x:c>`, `{"ref":"Y39494","type":"s","value":"100","inline_string":{"string_value":""}}`},
	{`<x:c r="BC41819" s="0"><x:v>0.00010642</x:v></x:c>`, `{"ref":"BC41819","type":"","value":"0.00010642","inline_string":{"string_value":""}}`},
	{`<x:c r="E47314" s="0" t="s"><x:v>807</x:v></x:c>`, `{"ref":"E47314","type":"s","value":"807","inline_string":{"string_value":""}}`},
	{`<x:c r="J51709" s="3" t="s"><x:v>64</x:v></x:c>`, `{"ref":"J51709","type":"s","value":"64","inline_string":{"string_value":""}}`},
	{`<x:c r="J56092" s="3" t="s"><x:v>78</x:v></x:c>`, `{"ref":"J56092","type":"s","value":"78","inline_string":{"string_value":""}}`},
	{`<x:c r="BF57189" s="0" t="s"><x:v>80</x:v></x:c>`, `{"ref":"BF57189","type":"s","value":"80","inline_string":{"string_value":""}}`},
	{`<x:c r="C59176" s="0" t="s"><x:v>193</x:v></x:c>`, `{"ref":"C59176","type":"s","value":"193","inline_string":{"string_value":""}}`},
	{`<x:c r="BC24056" s="0"><x:v>0.00054610</x:v></x:c>`, `{"ref":"BC24056","type":"","value":"0.00054610","inline_string":{"string_value":""}}`},
	{`<x:c r="C59849" s="0" t="s"><x:v>94</x:v></x:c>`, `{"ref":"C59849","type":"s","value":"94","inline_string":{"string_value":""}}`},
	{`<x:c r="AB61133" s="0"><x:v>1</x:v></x:c>`, `{"ref":"AB61133","type":"","value":"1","inline_string":{"string_value":""}}`},
	{`<x:c r="AF62437" s="0" t="s"><x:v>71</x:v></x:c>`, `{"ref":"AF62437","type":"s","value":"71","inline_string":{"string_value":""}}`},
	{`<x:c r="D70257" s="0" t="s"><x:v>97</x:v></x:c>`, `{"ref":"D70257","type":"s","value":"97","inline_string":{"string_value":""}}`},
	{`<x:c r="AY45007" s="0" t="s"><x:v>43</x:v></x:c>`, `{"ref":"AY45007","type":"s","value":"43","inline_string":{"string_value":""}}`},
	{`<x:c r="AR76607" s="0"><x:v>0.0005</x:v></x:c>`, `{"ref":"AR76607","type":"","value":"0.0005","inline_string":{"string_value":""}}`},
	{`<x:c r="AC54503" s="0"><x:v>0.00031391</x:v></x:c>`, `{"ref":"AC54503","type":"","value":"0.00031391","inline_string":{"string_value":""}}`},
	{`<x:c r="AE79091" s="0"><x:v>0.000750</x:v></x:c>`, `{"ref":"AE79091","type":"","value":"0.000750","inline_string":{"string_value":""}}`},
	{`<x:c r="E79463" s="0" t="s"><x:v>477</x:v></x:c>`, `{"ref":"E79463","type":"s","value":"477","inline_string":{"string_value":""}}`},
	{`<x:c r="F80333" s="0"><x:v>134</x:v></x:c>`, `{"ref":"F80333","type":"","value":"134","inline_string":{"string_value":""}}`},
	{`<x:c r="AE85474" s="0"><x:v>0.0004</x:v></x:c>`, `{"ref":"AE85474","type":"","value":"0.0004","inline_string":{"string_value":""}}`},
	{`<x:c r="M87565" s="0" t="s"><x:v>83</x:v></x:c>`, `{"ref":"M87565","type":"s","value":"83","inline_string":{"string_value":""}}`},
	{`<x:c r="AO88891" s="0"><x:v>0</x:v></x:c>`, `{"ref":"AO88891","type":"","value":"0","inline_string":{"string_value":""}}`},
	{`<x:c r="W89989" s="0"><x:v>0.0012</x:v></x:c>`, `{"ref":"W89989","type":"","value":"0.0012","inline_string":{"string_value":""}}`},
	{`<x:c r="W91853" s="0"><x:v>0.001278</x:v></x:c>`, `{"ref":"W91853","type":"","value":"0.001278","inline_string":{"string_value":""}}`},
	{`<x:c r="C93783" s="0" t="s"><x:v>272</x:v></x:c>`, `{"ref":"C93783","type":"s","value":"272","inline_string":{"string_value":""}}`},
	{`<x:c r="S94723" s="0"><x:v>3170</x:v></x:c>`, `{"ref":"S94723","type":"","value":"3170","inline_string":{"string_value":""}}`},
	{`<x:c r="Y98062" s="0" t="s"><x:v>100</x:v></x:c>`, `{"ref":"Y98062","type":"s","value":"100","inline_string":{"string_value":""}}`},
	{`<x:c r="AK99449" s="0"><x:v>0.000340</x:v></x:c>`, `{"ref":"AK99449","type":"","value":"0.000340","inline_string":{"string_value":""}}`},
	{`<x:c r="AD100307" s="0"><x:v>0.00050909</x:v></x:c>`, `{"ref":"AD100307","type":"","value":"0.00050909","inline_string":{"string_value":""}}`},
	{`<x:c r="BH104391" s="0"><x:v>0.000947</x:v></x:c>`, `{"ref":"BH104391","type":"","value":"0.000947","inline_string":{"string_value":""}}`},
	{`<x:c r="C106141" s="0" t="s"><x:v>104</x:v></x:c>`, `{"ref":"C106141","type":"s","value":"104","inline_string":{"string_value":""}}`},
	{`<x:c r="D106622" s="0" t="s"><x:v>115</x:v></x:c>`, `{"ref":"D106622","type":"s","value":"115","inline_string":{"string_value":""}}`},
	{`<x:c r="O109481" s="0" t="s"><x:v>68</x:v></x:c>`, `{"ref":"O109481","type":"s","value":"68","inline_string":{"string_value":""}}`},
	{`<x:c r="F109065" s="0"><x:v>520</x:v></x:c>`, `{"ref":"F109065","type":"","value":"520","inline_string":{"string_value":""}}`},
	{`<x:c r="BA110234" s="0" t="s"><x:v>74</x:v></x:c>`, `{"ref":"BA110234","type":"s","value":"74","inline_string":{"string_value":""}}`},
	{`<x:c r="AU78355" s="0" t="s"><x:v>72</x:v></x:c>`, `{"ref":"AU78355","type":"s","value":"72","inline_string":{"string_value":""}}`},
	{`<x:c r="AZ114241" s="0" t="s"><x:v>73</x:v></x:c>`, `{"ref":"AZ114241","type":"s","value":"73","inline_string":{"string_value":""}}`},
	{`<x:c r="I114752" s="0" t="s"><x:v>229</x:v></x:c>`, `{"ref":"I114752","type":"s","value":"229","inline_string":{"string_value":""}}`},
	{`<x:c r="H116960" s="0" t="s"><x:v>123</x:v></x:c>`, `{"ref":"H116960","type":"s","value":"123","inline_string":{"string_value":""}}`},
	{`<x:c r="D117883" s="0" t="s"><x:v>61</x:v></x:c>`, `{"ref":"D117883","type":"s","value":"61","inline_string":{"string_value":""}}`},
	{`<x:c r="W119153" s="0"><x:v>0.000547</x:v></x:c>`, `{"ref":"W119153","type":"","value":"0.000547","inline_string":{"string_value":""}}`},
	{`<x:c r="AT121330" s="0"><x:v>0.000980</x:v></x:c>`, `{"ref":"AT121330","type":"","value":"0.000980","inline_string":{"string_value":""}}`},
	{`<x:c r="E122284" s="0" t="s"><x:v>552</x:v></x:c>`, `{"ref":"E122284","type":"s","value":"552","inline_string":{"string_value":""}}`},
	{`<x:c r="BF94031" s="0" t="s"><x:v>80</x:v></x:c>`, `{"ref":"BF94031","type":"s","value":"80","inline_string":{"string_value":""}}`},
	{`<x:c r="K125014" s="0" t="s"><x:v>1457</x:v></x:c>`, `{"ref":"K125014","type":"s","value":"1457","inline_string":{"string_value":""}}`},
	{`<x:c r="AE126211" s="0"><x:v>0.0013</x:v></x:c>`, `{"ref":"AE126211","type":"","value":"0.0013","inline_string":{"string_value":""}}`},
	{`<x:c r="Y127892" s="0" t="s"><x:v>110</x:v></x:c>`, `{"ref":"Y127892","type":"s","value":"110","inline_string":{"string_value":""}}`},
	{`<x:c r="C132370" s="0" t="s"><x:v>111</x:v></x:c>`, `{"ref":"C132370","type":"s","value":"111","inline_string":{"string_value":""}}`},
	{`<x:c r="AC145194" s="0"><x:v>0.00027876</x:v></x:c>`, `{"ref":"AC145194","type":"","value":"0.00027876","inline_string":{"string_value":""}}`},
	{`<x:c r="BD145870" s="0"><x:v>0.002776</x:v></x:c>`, `{"ref":"BD145870","type":"","value":"0.002776","inline_string":{"string_value":""}}`},
	{`<x:c r="AT154103" s="0"><x:v>0.001130</x:v></x:c>`, `{"ref":"AT154103","type":"","value":"0.001130","inline_string":{"string_value":""}}`},
	{`<x:c r="AU154309" s="0" t="s"><x:v>72</x:v></x:c>`, `{"ref":"AU154309","type":"s","value":"72","inline_string":{"string_value":""}}`},
	{`<x:c r="AG156314" s="0"><x:v>0</x:v></x:c>`, `{"ref":"AG156314","type":"","value":"0","inline_string":{"string_value":""}}`},
	{`<x:c r="B158413" s="0" t="s"><x:v>60</x:v></x:c>`, `{"ref":"B158413","type":"s","value":"60","inline_string":{"string_value":""}}`},
	{`<x:c r="AY136265" s="0" t="s"><x:v>44</x:v></x:c>`, `{"ref":"AY136265","type":"s","value":"44","inline_string":{"string_value":""}}`},
	{`<x:c r="AK165552" s="0"><x:v>0.000310</x:v></x:c>`, `{"ref":"AK165552","type":"","value":"0.000310","inline_string":{"string_value":""}}`},
	{`<x:c r="BH165890" s="0"><x:v>0.0007</x:v></x:c>`, `{"ref":"BH165890","type":"","value":"0.0007","inline_string":{"string_value":""}}`},
	{`<x:c r="H141785" s="0" t="s"><x:v>184</x:v></x:c>`, `{"ref":"H141785","type":"s","value":"184","inline_string":{"string_value":""}}`},
	{`<x:c r="G170942" s="3" t="s"><x:v>3685</x:v></x:c>`, `{"ref":"G170942","type":"s","value":"3685","inline_string":{"string_value":""}}`},
	{`<x:c r="J116317" s="3" t="s"><x:v>78</x:v></x:c>`, `{"ref":"J116317","type":"s","value":"78","inline_string":{"string_value":""}}`},
	{`<x:c r="AZ176712" s="0" t="s"><x:v>73</x:v></x:c>`, `{"ref":"AZ176712","type":"s","value":"73","inline_string":{"string_value":""}}`},
	{`<x:c r="AF178276" s="0" t="s"><x:v>71</x:v></x:c>`, `{"ref":"AF178276","type":"s","value":"71","inline_string":{"string_value":""}}`},
	{`<x:c r="U148354" s="0"><x:v>1.67971612</x:v></x:c>`, `{"ref":"U148354","type":"","value":"1.67971612","inline_string":{"string_value":""}}`},
	{`<x:c r="B179677" s="0" t="s"><x:v>60</x:v></x:c>`, `{"ref":"B179677","type":"s","value":"60","inline_string":{"string_value":""}}`},
	{`<x:c r="O150287" s="0" t="s"><x:v>68</x:v></x:c>`, `{"ref":"O150287","type":"s","value":"68","inline_string":{"string_value":""}}`},
	{`<x:c r="AP180753" s="0"><x:v>0.9999</x:v></x:c>`, `{"ref":"AP180753","type":"","value":"0.9999","inline_string":{"string_value":""}}`},
	{`<x:c r="S151283" s="0"><x:v>2171</x:v></x:c>`, `{"ref":"S151283","type":"","value":"2171","inline_string":{"string_value":""}}`},
	{`<x:c r="AJ183605" s="0"><x:v>0.00038095</x:v></x:c>`, `{"ref":"AJ183605","type":"","value":"0.00038095","inline_string":{"string_value":""}}`},
	{`<x:c r="P184238" s="0" t="s"><x:v>69</x:v></x:c>`, `{"ref":"P184238","type":"s","value":"69","inline_string":{"string_value":""}}`},
	{`<x:c r="BH185822" s="0"><x:v>0.00002926</x:v></x:c>`, `{"ref":"BH185822","type":"","value":"0.00002926","inline_string":{"string_value":""}}`},
	{`<x:c r="D186592" s="0" t="s"><x:v>181</x:v></x:c>`, `{"ref":"D186592","type":"s","value":"181","inline_string":{"string_value":""}}`},
	{`<x:c r="X188717" s="0"><x:v>0.0011</x:v></x:c>`, `{"ref":"X188717","type":"","value":"0.0011","inline_string":{"string_value":""}}`},
	{`<x:c r="AO188874" s="0"><x:v>0</x:v></x:c>`, `{"ref":"AO188874","type":"","value":"0","inline_string":{"string_value":""}}`},
	{`<x:c r="N190954" s="0"><x:v>8</x:v></x:c>`, `{"ref":"N190954","type":"","value":"8","inline_string":{"string_value":""}}`},
	{`<x:c r="D191746" s="0" t="s"><x:v>94</x:v></x:c>`, `{"ref":"D191746","type":"s","value":"94","inline_string":{"string_value":""}}`},
	{`<x:c r="Y195084" s="0" t="s"><x:v>96</x:v></x:c>`, `{"ref":"Y195084","type":"s","value":"96","inline_string":{"string_value":""}}`},
	{`<x:c r="AG198133" s="0"><x:v>0</x:v></x:c>`, `{"ref":"AG198133","type":"","value":"0","inline_string":{"string_value":""}}`},
	{`<x:c r="M198475" s="0" t="s"><x:v>83</x:v></x:c>`, `{"ref":"M198475","type":"s","value":"83","inline_string":{"string_value":""}}`},
	{`<x:c r="AK201898" s="0"><x:v>0.001010</x:v></x:c>`, `{"ref":"AK201898","type":"","value":"0.001010","inline_string":{"string_value":""}}`},
	{`<x:c r="BA204585" s="0" t="s"><x:v>74</x:v></x:c>`, `{"ref":"BA204585","type":"s","value":"74","inline_string":{"string_value":""}}`},
	{`<x:c r="M205768" s="0" t="s"><x:v>67</x:v></x:c>`, `{"ref":"M205768","type":"s","value":"67","inline_string":{"string_value":""}}`},
	{`<x:c r="AS205805" s="0"><x:v>0.0010</x:v></x:c>`, `{"ref":"AS205805","type":"","value":"0.0010","inline_string":{"string_value":""}}`},
	{`<x:c r="AT207353" s="0"><x:v>0.000250</x:v></x:c>`, `{"ref":"AT207353","type":"","value":"0.000250","inline_string":{"string_value":""}}`},
	{`<x:c r="M208342" s="0" t="s"><x:v>67</x:v></x:c>`, `{"ref":"M208342","type":"s","value":"67","inline_string":{"string_value":""}}`},
	{`<x:c r="AE211397" s="0"><x:v>0.00131250</x:v></x:c>`, `{"ref":"AE211397","type":"","value":"0.00131250","inline_string":{"string_value":""}}`},
	{`<x:c r="O196250" s="0" t="s"><x:v>68</x:v></x:c>`, `{"ref":"O196250","type":"s","value":"68","inline_string":{"string_value":""}}`},
	{`<x:c r="E212483" s="0" t="s"><x:v>899</x:v></x:c>`, `{"ref":"E212483","type":"s","value":"899","inline_string":{"string_value":""}}`},
	{`<x:c r="AG213034" s="0"><x:v>0</x:v></x:c>`, `{"ref":"AG213034","type":"","value":"0","inline_string":{"string_value":""}}`},
	{`<x:c r="AB214952" s="0"><x:v>1</x:v></x:c>`, `{"ref":"AB214952","type":"","value":"1","inline_string":{"string_value":""}}`},
	{`<x:c r="AK218687" s="0"><x:v>0.000390</x:v></x:c>`, `{"ref":"AK218687","type":"","value":"0.000390","inline_string":{"string_value":""}}`},
	{`<x:c r="W206970" s="0"><x:v>0.001687</x:v></x:c>`, `{"ref":"W206970","type":"","value":"0.001687","inline_string":{"string_value":""}}`},
	{`<x:c r="G207042" s="3"><x:v>6664</x:v></x:c>`, `{"ref":"G207042","type":"","value":"6664","inline_string":{"string_value":""}}`},
	{`<x:c r="O216723" s="0" t="s"><x:v>68</x:v></x:c>`, `{"ref":"O216723","type":"s","value":"68","inline_string":{"string_value":""}}`},
	{`<x:c r="AP229210" s="0"><x:v>0.9999</x:v></x:c>`, `{"ref":"AP229210","type":"","value":"0.9999","inline_string":{"string_value":""}}`},
	{`<x:c r="P216963" s="0" t="s"><x:v>69</x:v></x:c>`, `{"ref":"P216963","type":"s","value":"69","inline_string":{"string_value":""}}`},
	{`<x:c r="H218111" s="0" t="s"><x:v>123</x:v></x:c>`, `{"ref":"H218111","type":"s","value":"123","inline_string":{"string_value":""}}`},
	{`<x:c r="BH219400" s="0"><x:v>0.00100933</x:v></x:c>`, `{"ref":"BH219400","type":"","value":"0.00100933","inline_string":{"string_value":""}}`},
	{`<x:c r="K233998" s="0" t="s"><x:v>793</x:v></x:c>`, `{"ref":"K233998","type":"s","value":"793","inline_string":{"string_value":""}}`},
	{`<x:c r="D241837" s="0" t="s"><x:v>181</x:v></x:c>`, `{"ref":"D241837","type":"s","value":"181","inline_string":{"string_value":""}}`},
	{`<x:c r="BF242832" s="0" t="s"><x:v>80</x:v></x:c>`, `{"ref":"BF242832","type":"s","value":"80","inline_string":{"string_value":""}}`},
	{`<x:c r="AZ244525" s="0" t="s"><x:v>73</x:v></x:c>`, `{"ref":"AZ244525","type":"s","value":"73","inline_string":{"string_value":""}}`},
	{`<x:c r="AK247967" s="0"><x:v>0.000010</x:v></x:c>`, `{"ref":"AK247967","type":"","value":"0.000010","inline_string":{"string_value":""}}`},
	{`<x:c r="AP248821" s="0"><x:v>0.9999</x:v></x:c>`, `{"ref":"AP248821","type":"","value":"0.9999","inline_string":{"string_value":""}}`},
	{`<x:c r="C237720" s="0" t="s"><x:v>109</x:v></x:c>`, `{"ref":"C237720","type":"s","value":"109","inline_string":{"string_value":""}}`},
	{`<x:c r="AM237707" s="0"><x:v>0.00197297</x:v></x:c>`, `{"ref":"AM237707","type":"","value":"0.00197297","inline_string":{"string_value":""}}`},
	{`<x:c r="BH254725" s="0"><x:v>0.00130430</x:v></x:c>`, `{"ref":"BH254725","type":"","value":"0.00130430","inline_string":{"string_value":""}}`},
	{`<x:c r="J254970" s="3" t="s"><x:v>78</x:v></x:c>`, `{"ref":"J254970","type":"s","value":"78","inline_string":{"string_value":""}}`},
	{`<x:c r="AK257908" s="0"><x:v>0.000570</x:v></x:c>`, `{"ref":"AK257908","type":"","value":"0.000570","inline_string":{"string_value":""}}`},
	{`<x:c r="AT258918" s="0"><x:v>0.0011</x:v></x:c>`, `{"ref":"AT258918","type":"","value":"0.0011","inline_string":{"string_value":""}}`},
	{`<x:c r="BF267064" s="0" t="s"><x:v>75</x:v></x:c>`, `{"ref":"BF267064","type":"s","value":"75","inline_string":{"string_value":""}}`},
	{`<x:c r="Y248039" s="0" t="s"><x:v>70</x:v></x:c>`, `{"ref":"Y248039","type":"s","value":"70","inline_string":{"string_value":""}}`},
	{`<x:c r="X272571" s="0"><x:v>0.00006851</x:v></x:c>`, `{"ref":"X272571","type":"","value":"0.00006851","inline_string":{"string_value":""}}`},
	{`<x:c r="O272924" s="0" t="s"><x:v>68</x:v></x:c>`, `{"ref":"O272924","type":"s","value":"68","inline_string":{"string_value":""}}`},
	{`<x:c r="AE274542" s="0"><x:v>0.00207637</x:v></x:c>`, `{"ref":"AE274542","type":"","value":"0.00207637","inline_string":{"string_value":""}}`},
	{`<x:c r="X253981" s="0"><x:v>0.00063836</x:v></x:c>`, `{"ref":"X253981","type":"","value":"0.00063836","inline_string":{"string_value":""}}`},
	{`<x:c r="I280252" s="0" t="s"><x:v>63</x:v></x:c>`, `{"ref":"I280252","type":"s","value":"63","inline_string":{"string_value":""}}`},
	{`<x:c r="BA280831" s="0" t="s"><x:v>74</x:v></x:c>`, `{"ref":"BA280831","type":"s","value":"74","inline_string":{"string_value":""}}`},
	{`<x:c r="AM282790" s="0"><x:v>0.00338221</x:v></x:c>`, `{"ref":"AM282790","type":"","value":"0.00338221","inline_string":{"string_value":""}}`},
	{`<x:c r="AY284643" s="0" t="s"><x:v>85</x:v></x:c>`, `{"ref":"AY284643","type":"s","value":"85","inline_string":{"string_value":""}}`},
	{`<x:c r="AT286751" s="0"><x:v>0.0020</x:v></x:c>`, `{"ref":"AT286751","type":"","value":"0.0020","inline_string":{"string_value":""}}`},
	{`<x:c r="BG288136" s="0"><x:v>0.00085280</x:v></x:c>`, `{"ref":"BG288136","type":"","value":"0.00085280","inline_string":{"string_value":""}}`},
	{`<x:c r="AE294533" s="0"><x:v>0.0003</x:v></x:c>`, `{"ref":"AE294533","type":"","value":"0.0003","inline_string":{"string_value":""}}`},
	{`<x:c r="AM295088" s="0"><x:v>0.00151301</x:v></x:c>`, `{"ref":"AM295088","type":"","value":"0.00151301","inline_string":{"string_value":""}}`},
	{`<x:c r="AL296287" s="0"><x:v>0.001170</x:v></x:c>`, `{"ref":"AL296287","type":"","value":"0.001170","inline_string":{"string_value":""}}`},
	{`<x:c r="BE298912" s="0"><x:v>0</x:v></x:c>`, `{"ref":"BE298912","type":"","value":"0","inline_string":{"string_value":""}}`},
	{`<x:c r="AZ284142" s="0" t="s"><x:v>73</x:v></x:c>`, `{"ref":"AZ284142","type":"s","value":"73","inline_string":{"string_value":""}}`},
	{`<x:c r="E304821" s="0" t="s"><x:v>417</x:v></x:c>`, `{"ref":"E304821","type":"s","value":"417","inline_string":{"string_value":""}}`},
	{`<x:c r="AJ307232" s="0"><x:v>0.00020021</x:v></x:c>`, `{"ref":"AJ307232","type":"","value":"0.00020021","inline_string":{"string_value":""}}`},
	{`<x:c r="AB308153" s="0"><x:v>0</x:v></x:c>`, `{"ref":"AB308153","type":"","value":"0","inline_string":{"string_value":""}}`},
	{`<x:c r="BD312240" s="0"><x:v>0.001369</x:v></x:c>`, `{"ref":"BD312240","type":"","value":"0.001369","inline_string":{"string_value":""}}`},
	{`<x:c r="I313146" s="0" t="s"><x:v>663</x:v></x:c>`, `{"ref":"I313146","type":"s","value":"663","inline_string":{"string_value":""}}`},
	{`<x:c r="AM313735" s="0"><x:v>0.00210534</x:v></x:c>`, `{"ref":"AM313735","type":"","value":"0.00210534","inline_string":{"string_value":""}}`},
	{`<x:c r="V313737" s="0"><x:v>0</x:v></x:c>`, `{"ref":"V313737","type":"","value":"0","inline_string":{"string_value":""}}`},
	{`<x:c r="AD294161" s="0"><x:v>0.00137083</x:v></x:c>`, `{"ref":"AD294161","type":"","value":"0.00137083","inline_string":{"string_value":""}}`},
	{`<x:c r="M317674" s="0" t="s"><x:v>67</x:v></x:c>`, `{"ref":"M317674","type":"s","value":"67","inline_string":{"string_value":""}}`},
	{`<x:c r="AK311078" s="0"><x:v>0.000310</x:v></x:c>`, `{"ref":"AK311078","type":"","value":"0.000310","inline_string":{"string_value":""}}`},
	{`<x:c r="A311689" s="2"><x:v>704372</x:v></x:c>`, `{"ref":"A311689","type":"","value":"704372","inline_string":{"string_value":""}}`},
	{`<x:c r="BD317104" s="0"><x:v>0.00046964</x:v></x:c>`, `{"ref":"BD317104","type":"","value":"0.00046964","inline_string":{"string_value":""}}`},
	{`<x:c r="M327084" s="0" t="s"><x:v>67</x:v></x:c>`, `{"ref":"M327084","type":"s","value":"67","inline_string":{"string_value":""}}`},
	{`<x:c r="AF328722" s="0" t="s"><x:v>71</x:v></x:c>`, `{"ref":"AF328722","type":"s","value":"71","inline_string":{"string_value":""}}`},
	{`<x:c r="AM331041" s="0"><x:v>0.00232625</x:v></x:c>`, `{"ref":"AM331041","type":"","value":"0.00232625","inline_string":{"string_value":""}}`},
	{`<x:c r="M332280" s="0" t="s"><x:v>67</x:v></x:c>`, `{"ref":"M332280","type":"s","value":"67","inline_string":{"string_value":""}}`},
	{`<x:c r="AB332970" s="0"><x:v>1</x:v></x:c>`, `{"ref":"AB332970","type":"","value":"1","inline_string":{"string_value":""}}`},
	{`<x:c r="I335288" s="0" t="s"><x:v>229</x:v></x:c>`, `{"ref":"I335288","type":"s","value":"229","inline_string":{"string_value":""}}`},
	{`<x:c r="D324580" s="0" t="s"><x:v>61</x:v></x:c>`, `{"ref":"D324580","type":"s","value":"61","inline_string":{"string_value":""}}`},
	{`<x:c r="AF324922" s="0" t="s"><x:v>71</x:v></x:c>`, `{"ref":"AF324922","type":"s","value":"71","inline_string":{"string_value":""}}`},
	{`<x:c r="AO327192" s="0"><x:v>0</x:v></x:c>`, `{"ref":"AO327192","type":"","value":"0","inline_string":{"string_value":""}}`},
	{`<x:c r="BD327856" s="0"><x:v>0.002776</x:v></x:c>`, `{"ref":"BD327856","type":"","value":"0.002776","inline_string":{"string_value":""}}`},
	{`<x:c r="H341254" s="0" t="s"><x:v>123</x:v></x:c>`, `{"ref":"H341254","type":"s","value":"123","inline_string":{"string_value":""}}`},
	{`<x:c r="X333015" s="0"><x:v>0.00090738</x:v></x:c>`, `{"ref":"X333015","type":"","value":"0.00090738","inline_string":{"string_value":""}}`},
	{`<x:c r="H341106" s="0" t="s"><x:v>123</x:v></x:c>`, `{"ref":"H341106","type":"s","value":"123","inline_string":{"string_value":""}}`},
	{`<x:c r="K349104" s="0" t="s"><x:v>3531</x:v></x:c>`, `{"ref":"K349104","type":"s","value":"3531","inline_string":{"string_value":""}}`},
	{`<x:c r="D344328" s="0" t="s"><x:v>61</x:v></x:c>`, `{"ref":"D344328","type":"s","value":"61","inline_string":{"string_value":""}}`},
	{`<x:c r="BG346902" s="0"><x:v>0.000655</x:v></x:c>`, `{"ref":"BG346902","type":"","value":"0.000655","inline_string":{"string_value":""}}`},
	{`<x:c r="M347970" s="0" t="s"><x:v>67</x:v></x:c>`, `{"ref":"M347970","type":"s","value":"67","inline_string":{"string_value":""}}`},
	{`<x:c r="AP357189" s="0"><x:v>0.9999</x:v></x:c>`, `{"ref":"AP357189","type":"","value":"0.9999","inline_string":{"string_value":""}}`},
	{`<x:c r="E357245" s="0" t="s"><x:v>182</x:v></x:c>`, `{"ref":"E357245","type":"s","value":"182","inline_string":{"string_value":""}}`},
	{`<x:c r="M359415" s="0" t="s"><x:v>67</x:v></x:c>`, `{"ref":"M359415","type":"s","value":"67","inline_string":{"string_value":""}}`},
	{`<x:c r="BG362419" s="0"><x:v>0.00081450</x:v></x:c>`, `{"ref":"BG362419","type":"","value":"0.00081450","inline_string":{"string_value":""}}`},
	{`<x:c r="AN364339" s="0"><x:v>0.0003</x:v></x:c>`, `{"ref":"AN364339","type":"","value":"0.0003","inline_string":{"string_value":""}}`},
	{`<x:c r="F367668" s="0"><x:v>730</x:v></x:c>`, `{"ref":"F367668","type":"","value":"730","inline_string":{"string_value":""}}`},
	{`<x:c r="BH368815" s="0"><x:v>0.01543458</x:v></x:c>`, `{"ref":"BH368815","type":"","value":"0.01543458","inline_string":{"string_value":""}}`},
	{`<x:c r="M368703" s="0" t="s"><x:v>83</x:v></x:c>`, `{"ref":"M368703","type":"s","value":"83","inline_string":{"string_value":""}}`},
	{`<x:c r="O375589" s="0" t="s"><x:v>68</x:v></x:c>`, `{"ref":"O375589","type":"s","value":"68","inline_string":{"string_value":""}}`},
	{`<x:c r="O376437" s="0" t="s"><x:v>68</x:v></x:c>`, `{"ref":"O376437","type":"s","value":"68","inline_string":{"string_value":""}}`},
	{`<x:c r="AO384698" s="0"><x:v>0</x:v></x:c>`, `{"ref":"AO384698","type":"","value":"0","inline_string":{"string_value":""}}`},
	{`<x:c r="AD385061" s="0"><x:v>0.00217816</x:v></x:c>`, `{"ref":"AD385061","type":"","value":"0.00217816","inline_string":{"string_value":""}}`},
	{`<c r="E4195" s="10"><v>0.00141</v></c>`, `{"ref":"E4195","type":"","value":"0.00141","inline_string":{"string_value":""}}`},
	{`<c r="B21777" s="8"><v>909432</v></c>`, `{"ref":"B21777","type":"","value":"909432","inline_string":{"string_value":""}}`},
	{`<c r="G29381" s="11"><v>44.68333333333334</v></c>`, `{"ref":"G29381","type":"","value":"44.68333333333334","inline_string":{"string_value":""}}`},
	{`<c r="H43002" s="10"><v>0.07692500000000001</v></c>`, `{"ref":"H43002","type":"","value":"0.07692500000000001","inline_string":{"string_value":""}}`},
	{`<c r="H47934" s="10"><v>0</v></c>`, `{"ref":"H47934","type":"","value":"0","inline_string":{"string_value":""}}`},
	{`<c r="D50887" s="9" t="inlineStr"><is><t>06/12/18 (Tue)</t></is></c>`, `{"ref":"D50887","type":"inlineStr","value":"","inline_string":{"string_value":"06/12/18 (Tue)"}}`},
	{`<c r="H64672" s="10"><v>0</v></c>`, `{"ref":"H64672","type":"","value":"0","inline_string":{"string_value":""}}`},
	{`<c r="D66220" s="9" t="inlineStr"><is><t>06/12/18 (Tue)</t></is></c>`, `{"ref":"D66220","type":"inlineStr","value":"","inline_string":{"string_value":"06/12/18 (Tue)"}}`},
	{`<c r="B69714" s="8"><v>910739</v></c>`, `{"ref":"B69714","type":"","value":"910739","inline_string":{"string_value":""}}`},
	{`<c r="E84949" s="10"><v>0.00138</v></c>`, `{"ref":"E84949","type":"","value":"0.00138","inline_string":{"string_value":""}}`},
	{`<c r="D89508" s="9" t="inlineStr"><is><t>06/11/18 (Mon)</t></is></c>`, `{"ref":"D89508","type":"inlineStr","value":"","inline_string":{"string_value":"06/11/18 (Mon)"}}`},
	{`<c r="H94935" s="10"><v>0</v></c>`, `{"ref":"H94935","type":"","value":"0","inline_string":{"string_value":""}}`},
	{`<c r="A107666" s="7" t="inlineStr"><is><t>Everyman</t></is></c>`, `{"ref":"A107666","type":"inlineStr","value":"","inline_string":{"string_value":"Everyman"}}`},
	{`<c r="A108189" s="7" t="inlineStr"><is><t>Everyman</t></is></c>`, `{"ref":"A108189","type":"inlineStr","value":"","inline_string":{"string_value":"Everyman"}}`},
	{`<c r="D108491" s="9" t="inlineStr"><is><t>06/12/18 (Tue)</t></is></c>`, `{"ref":"D108491","type":"inlineStr","value":"","inline_string":{"string_value":"06/12/18 (Tue)"}}`},
	{`<c r="H138134" s="10"><v>0.00219</v></c>`, `{"ref":"H138134","type":"","value":"0.00219","inline_string":{"string_value":""}}`},
	{`<c r="F147683" s="10"><v>0.001575</v></c>`, `{"ref":"F147683","type":"","value":"0.001575","inline_string":{"string_value":""}}`},
	{`<c r="G174390" s="11"><v>0</v></c>`, `{"ref":"G174390","type":"","value":"0","inline_string":{"string_value":""}}`},
	{`<c r="G179852" s="11"><v>0.1666666666666667</v></c>`, `{"ref":"G179852","type":"","value":"0.1666666666666667","inline_string":{"string_value":""}}`},
	{`<c r="F189829" s="10"><v>0.003375</v></c>`, `{"ref":"F189829","type":"","value":"0.003375","inline_string":{"string_value":""}}`},
	{`<c r="G198514" s="11"><v>249.5333333333333</v></c>`, `{"ref":"G198514","type":"","value":"249.5333333333333","inline_string":{"string_value":""}}`},
	{`<c r="A206073" s="7" t="inlineStr"><is><t>Everyman</t></is></c>`, `{"ref":"A206073","type":"inlineStr","value":"","inline_string":{"string_value":"Everyman"}}`},
	{`<c r="G213378" s="11"><v>0.2</v></c>`, `{"ref":"G213378","type":"","value":"0.2","inline_string":{"string_value":""}}`},
	{`<c r="H218746" s="10"><v>0.000136</v></c>`, `{"ref":"H218746","type":"","value":"0.000136","inline_string":{"string_value":""}}`},
	{`<c r="A231355" s="7" t="inlineStr"><is><t>Everyman</t></is></c>`, `{"ref":"A231355","type":"inlineStr","value":"","inline_string":{"string_value":"Everyman"}}`},
	{`<c r="D233738" s="9" t="inlineStr"><is><t>06/13/18 (Wed)</t></is></c>`, `{"ref":"D233738","type":"inlineStr","value":"","inline_string":{"string_value":"06/13/18 (Wed)"}}`},
	{`<c r="H235298" s="10"><v>0.009510000000000005</v></c>`, `{"ref":"H235298","type":"","value":"0.009510000000000005","inline_string":{"string_value":""}}`},
	{`<c r="H246200" s="10"><v>0.003359</v></c>`, `{"ref":"H246200","type":"","value":"0.003359","inline_string":{"string_value":""}}`},
	{`<c r="B256492" s="8"><v>205497</v></c>`, `{"ref":"B256492","type":"","value":"205497","inline_string":{"string_value":""}}`},
	{`<c r="B270054" s="8"><v>817451</v></c>`, `{"ref":"B270054","type":"","value":"817451","inline_string":{"string_value":""}}`},
	{`<c r="C289153" s="9" t="inlineStr"><is><t>Hamburger</t></is></c>`, `{"ref":"C289153","type":"inlineStr","value":"","inline_string":{"string_value":"Hamburger"}}`},
	{`<c r="F302366" s="10"><v>0</v></c>`, `{"ref":"F302366","type":"","value":"0","inline_string":{"string_value":""}}`},
	{`<c r="D303117" s="9" t="inlineStr"><is><t>06/11/18 (Mon)</t></is></c>`, `{"ref":"D303117","type":"inlineStr","value":"","inline_string":{"string_value":"06/11/18 (Mon)"}}`},
	{`<c r="F312129" s="10"><v>0.000232</v></c>`, `{"ref":"F312129","type":"","value":"0.000232","inline_string":{"string_value":""}}`},
	{`<c r="C313697" s="9" t="inlineStr"><is><t>Eggs</t></is></c>`, `{"ref":"C313697","type":"inlineStr","value":"","inline_string":{"string_value":"Eggs"}}`},
	{`<c r="G317000" s="11"><v>0</v></c>`, `{"ref":"G317000","type":"","value":"0","inline_string":{"string_value":""}}`},
}
