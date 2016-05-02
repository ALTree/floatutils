package floats_test

import (
	"math"
	"math/big"
	"math/rand"
	"testing"

	"github.com/ALTree/floats"
)

func TestExp(t *testing.T) {
	for _, test := range []struct {
		x    string
		want string
	}{
		{"0", "1"},
		{"1", "2.7182818284590452353602874713526624977572470936999595749669676277240766303535475945713821785251664274274663919320030599218174135966290435729003342952605956307381323286279434907632338298807531952510190115738341879307021540891499348841675092447614606680822648001684774118537423454424371075390777449920695517027618386062613313845830007520449338265602976"},
		{"1.5", "4.4816890703380648226020554601192758190057498683696670567726500827859366744667137729810538313824533913886163506518301957689627464772204086069617596449736935381785298966216866555101351015468252930697652168582207145843214628870201383138594206299440366192845735003904511744769330306157776861818063974846024442278949433305735015582659896397844432731673273"},
		{"2", "7.3890560989306502272304274605750078131803155705518473240871278225225737960790577633843124850791217947737531612654788661238846036927812733744783922133980777749001228956074107537023913309475506820865818202696478682084042209822552348757424625414146799281293318880707633010193378997407299869600953033075153208188236846947930299135587714456831239232727646"},
		{"3", "20.085536923187667740928529654581717896987907838554150144378934229698845878091973731204497160253017702153607615851949002881811012479353506690232621784477250503945677100066077851812229047884383940258152534709352622981465538424555697733515108150118404754933838497843177676070913772862491787349396037822793717687131254060597553426640826030948663920216259"},

		{"-1", "0.36787944117144232159552377016146086744581113103176783450783680169746149574489980335714727434591964374662732527684399520824697579279012900862665358949409878309219436737733811504863899112514561634498771997868447595793974730254989249545323936620796481051464752061229422308916492656660036507457728370553285373838810680478761195682989345449735073931859922"},
		{"-2", "0.13533528323661269189399949497248440340763154590957588146815887265407337410148768993709812249065704875507728718963355221244934687189285303815889513499670600559125022755868258230483842057584538468003599408344602481287135375015664353399593608501390049529421705857601948571122397095990883595090571764528251279379538022237440385068269131295459365886804367"},
		{"-3", "0.049787068367863942979342415650061776631699592188423215567627727606060667730199550154054244236633344526401328650893681950864643386736174297123488422626590132549710257089250891729183705544267766471294627261313755158051249249208013335774449487985072339959233419058693861230319791977014791562486437888837044087835498513120050395020976909328170160274676519"},

		{"10", "22026.465794806716516957900645284244366353512618556781074235426355225202818570792575199120968164525895451555501092457836652423291606522895166222480137728972873485577837847275195480610095881417055888657927317236168401192698035170264925041101757502556764762696107543817931960834044404934236682455357614946828619042431465132389556031319229262768101604495"},
		{"100", "2.6881171418161354484126255515800135873611118773741922415191608615280287034909564914158871097219845710811670879190576068697597709761868233548459638929871966089629133626120029380957276534032962269865668016917743514451846065162804442237756762296960284731911402129862281040057911593878790384974173340084912432828126815454426051808828625966509400466909062e43"},
		{"1000", "1.9700711140170469938888793522433231253169379853238457899528029913850638507824411934749780765630268899309638179875202269359829817305446128992326278366015282523232053516958456675619227156760278807142246682631400685516850865349794166031604536781793809290529972858013286994585647028653437590045656435558915622042232026051882611228863835837224872472521451e434"},

		{"-10", "0.000045399929762484851535591515560550610237918088866564969259071305650999421614302281652525004545947782321708055089686028492945199117244520388837183347709414567560990909217007363970181059501783900762968517787030908824365171548448722293652332416020501168264360305604941570107729975354408079403994232932138270780520042710498960354486166066837009201707573209"},
		{"-100", "3.7200759760208359629596958038631183373588922923767819671206138766632904758958157181571187786422814966019356176423110698002479856420525356002661856882839075574388191160228448691497585855102816611741608772370701345082175755257496876380478927279529400619796226477050521097935092405571614981699373980650794385017392666116669084820355852767349264735965334e-44"},
		{"-1000", "5.0759588975494567652918094795743369193055992828928373618323938454105405429748191756796621690465428678636671068310652851135787934480190632251259072300213915638091771495398351108574919194309548129952421441572726108465407163812260104924530270737073247546217081943180823516857873407345613076984468096760005536701904004361380296144254899617340297251706670e-435"},
	} {
		for _, prec := range []uint{24, 53, 64, 100, 200, 300, 400, 500, 600, 700, 800, 900, 1000} {
			want := new(big.Float).SetPrec(prec)
			want.Parse(test.want, 10)

			x := new(big.Float).SetPrec(prec)
			x.Parse(test.x, 10)

			z := floats.Exp(x)

			if z.Cmp(want) != 0 {
				t.Errorf("prec = %d, Exp(%v) =\ngot  %g;\nwant %g", prec, test.x, z, want)
			}
		}
	}
}

func TestExp32(t *testing.T) {
	for i := 0; i < 1e4; i++ {
		r := rand.Float32() * 100
		x := big.NewFloat(float64(r)).SetPrec(24)
		z, acc := floats.Exp(x).Float32()
		want := math.Exp(float64(r))
		if z != float32(want) {
			t.Errorf("Exp(%f) =\n got %b (%s);\nwant %b (Exact)", x, z, acc, float32(want))
		}
	}
}

func TestExp32Small(t *testing.T) {
	for i := 0; i < 1e4; i++ {
		r := rand.Float32() * 1e-20
		x := big.NewFloat(float64(r)).SetPrec(24)
		z, acc := floats.Exp(x).Float32()
		want := math.Exp(float64(r))
		if z != float32(want) {
			t.Errorf("Exp(%f) =\n got %b (%s);\nwant %b (Exact)", x, z, acc, float32(want))
		}
	}
}

func TestExp64(t *testing.T) {
	for i := 0; i < 1e4; i++ {
		r := rand.Float64() * 100
		x := big.NewFloat(r).SetPrec(53)
		z, acc := floats.Exp(x).Float64()
		want := math.Exp(r)
		if math.Abs(z-want)/want > 1e-14 {
			t.Errorf("Exp(%g) =\n got %b (%s);\nwant %b (Exact)", x, z, acc, want)
		}
	}
}

func TestExp64Small(t *testing.T) {
	for i := 0; i < 1e4; i++ {
		r := rand.Float64() * 1e-20
		x := big.NewFloat(r).SetPrec(53)
		z, acc := floats.Exp(x).Float64()
		want := math.Exp(r)
		if math.Abs(z-want)/want > 1e-14 {
			t.Errorf("Exp(%g) =\n got %b (%s);\nwant %b (Exact)", x, z, acc, want)
		}
	}
}

// ---------- Benchmarks ----------

func benchmarkExp(num float64, prec uint, b *testing.B) {
	b.ReportAllocs()
	x := new(big.Float).SetPrec(prec).SetFloat64(num)
	for n := 0; n < b.N; n++ {
		floats.Exp(x)
	}
}

func BenchmarkExp2Prec53(b *testing.B)     { benchmarkExp(2, 53, b) }
func BenchmarkExp2Prec100(b *testing.B)    { benchmarkExp(2, 1e2, b) }
func BenchmarkExp2Prec1000(b *testing.B)   { benchmarkExp(2, 1e3, b) }
func BenchmarkExp2Prec10000(b *testing.B)  { benchmarkExp(2, 1e4, b) }
func BenchmarkExp2Prec100000(b *testing.B) { benchmarkExp(2, 1e5, b) }
