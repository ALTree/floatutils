package floats_test

import (
	"math/big"
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
