func main() {

	crea_perso().menu()
}

func (p Personnage) menu() {
	// scan --> switch case
	var retour string
	fmt.Println("1 : Accéder aux informations du personnage")
	fmt.Println("2 : Accéder au contenu de l’inventaire")
	fmt.Println("3 : Marchand")
	fmt.Println("4 : Combattre")
	fmt.Println("10 : Quitter (vous perdez votre sauvegarde et devez recréer un personnage).")
	fmt.Println("Que souhaitez vous faire (tapez le numéro du choix correspondant) : ")
	fmt.Scan(&retour)
	switch retour {
	case "1":
		p.displayinfo()
	case "2":
		p.accessInventory()
	case "3":
		var marchand Marchand
		marchand.displayinfo()
	case "4":
		p.creer_combat()
	case "10":
		crea_perso().menu()
	}
}

func (p *Personnage) Init(var1 string, var2 string, var3 int, var4 int, var5 int, var6 []item, var7 []competence, var8 equipement, var9 int, var10 int, var11 int, var12 int, var13 int) {
	p.nom_p = var1
	p.classe = var2
	p.niveau = var3
	p.pv_max = var4
	p.pv_now = var5
	p.inventaire = var6
	p.skill = var7
	p.stuff = var8
	p.credit = var9
	p.xp = var10
	p.mana_now = var11
	p.mana_max = var12
	p.initiative = var13
}

func (p Personnage) accessInventory() {
	for _, element := range p.inventaire {
		element.displayinfo()
	}
	if len(p.inventaire) == 0 {
		fmt.Println("Votre inventaire est vide.")
	}
}

func (p Personnage) is_dead() bool {
	if p.pv_now <= 0 {
		return true
	} else {
		return false
	}
}
func (p *Personnage) dead() {
	if p.is_dead() {
		fmt.Println("Vous êtes mort.\n Vous retrouvez la moitié de vous pv max et retourez au menu principal.")
		time.Sleep(2 * time.Second)
		p.pv_now = p.pv_max / 2

	}
}

func (p *Personnage) spellbook(comp competence) {
	test := false
	for _, element := range p.skill {
		if element.id == comp.id {
			test = true
		}
	}
	if !test {
		p.skill = append(p.skill, comp)
		fmt.Println("Félicitation, vous avez appris la compétence ", comp.nom)
	} else {
		fmt.Println("Vous connaissez déjà cette compétence.")
	}
}

type forge struct {
	list_craft []item
}

func (f forge) displayinfo() {
	fmt.Println("La forge propose :")
	for _, element := range f.list_craft {
		list_ressources := []string{}
		for _, element2 := range element.ressources {
			list_ressources = append(list_ressources, element2.nom)
		}
		fmt.Println(element.nom, "en utilisant", list_ressources, ".")
	}
}

func ToLower(s string) string {
	resultat := ""

	for i := 0; i < len(s); i++ {
		if s[i] < 91 && s[i] > 64 {
			resultat += string(rune(int(s[i]) + 32))
		} else {
			resultat += string(s[i])
		}
	}
	return resultat
}

func ToUpper(s string) string {
	resultat := ""

	for i := 0; i < len(s); i++ {
		if s[i] < 123 && s[i] > 96 {
			resultat += string(rune(int(s[i]) - 32))
		} else {
			resultat += string(s[i])
		}
	}
	return resultat
}
func Capitalize(s string) string {
	resultat := ToUpper(string(s[0]))
	for i := 1; i < len(s); i++ {
		if string(s[i]) == " " || string(s[i]) == "+" {
			resultat += string(s[i]) + ToUpper(string(s[i+1]))
			i += 1
		} else {
			resultat += ToLower(string(s[i]))
		}

	}
	return resultat
}

func IsAlpha(s string) bool {
	if s == "" {
		return false
	}

	for i := 0; i < len(s); i++ {
		if s[i] < 65 || s[i] > 91 {
			if s[i] < 96 || s[i] > 123 {
				return false
			}
		}

	}
	return true
}

type equipement struct {
	protection int
	helmet     item
	bodyshield item
	legshield  item
}

func (e *equipement) maj() {
	e.protection = e.helmet.protection + e.bodyshield.protection + e.legshield.protection
}
func (e equipement) pv() int {
	return e.protection
}

type monstre struct {
	id_monstre  int
	nom_m       string
	pv_max      int
	pv_now      int
	attaque     int
	initiative  int
	xp_donne    int
	tour_patern int
	recompense  []item
}

func creer_monstre(var1 int, var2 string, var3 int, var4 int, var5 int, var6 int, var7 int, var8 int, var9 []item) monstre {
	var m monstre
	m.id_monstre = var1
	m.nom_m = var2
	m.pv_max = var3
	m.pv_now = var4
	m.attaque = var5
	m.initiative = var6
	m.xp_donne = var7
	m.tour_patern = var8
	m.recompense = var9
	return m
}

type arene struct {
	joueur      Personnage
	ennemi      monstre
	info_joueur *Personnage
	info_ennemi *monstre
}

func (p Personnage) creer_combat() {
	lmonstre := []monstre{creer_monstre(1, "Gobelin", 40, 40, 5, 1, 10, 3, []item{}), creer_monstre(2, "Orc", 60, 60, 8, 3, 25, 4, []item{}), creer_monstre(3, "Troll", 90, 90, 15, 5, 50, 5, []item{})}
	var a arene
	var m monstre
	fmt.Println("Qui souhaitez vous affronter ?")
	for _, element := range lmonstre[:p.niveau] {
		fmt.Println(element.id_monstre, ":", element.nom_m, "qui vous donnera", element.xp_donne, "points d'expérience et", element.recompense, ".")
	}
	var choix int
	fmt.Scan(&choix)
	for _, element := range lmonstre {
		if element.id_monstre == choix {
			m = element
		}
	}
	a.Init(p, m)
	a.combattre()
}
func (a *arene) Init(j1 Personnage, j2 monstre) {
	a.joueur = j1
	a.ennemi = j2
	a.info_joueur = &a.joueur
	a.info_ennemi = &a.ennemi
}

func (a *arene) joueur_gagne() {
	fmt.Println("Félicitation, vous avez gagné. \n Votre joueur regagne la totalité de ses pv et vous retournez au menu principal.")
	fmt.Println("Le monstre que vous venez de battre vous rapporte", a.ennemi.xp_donne, "XP.")
	a.info_joueur.xp += a.ennemi.xp_donne
	reste := 0
	if a.joueur.xp >= a.joueur.niveau*100 {
		a.info_joueur.niveau++
		reste = (a.joueur.xp - (a.joueur.niveau * 100))
		a.info_joueur.xp = reste
		fmt.Println("Vous gagnez un niveau !!! \n Vous êtes maintenant niveau", a.joueur.niveau, ".")
		time.Sleep(2 * time.Second)
	}
	fmt.Println("Vous avez", a.joueur.xp, "/", a.joueur.xp, "XP.")
	time.Sleep(3 * time.Second)
	a.joueur.menu()
}

func (p *Personnage) attaquer(m *monstre) {
	fmt.Println("\nLes attaques à votre disposition sont les suivantes :")
	for i, element := range p.skill {
		fmt.Println(i, ":", element.nom, "qui inflige", element.degats, "dégats.")
	}
	fmt.Println("\nQuelle attaque choisissez vous ? \n (choisissez le numéro correspondant)")
	var capa int
	fmt.Scan(&capa)
	attaque := p.skill[capa].degats
	fmt.Println(p.nom_p, "attaque", m.nom_m, "et lui inflige", attaque, "dégats")
	m.pv_now -= attaque
	fmt.Println("points de vie actuel du monstre : ", m.pv_now, "/", m.pv_max)
}

func (a *arene) combattre() {
	fmt.Println("Vous êtes face à un", a.ennemi.nom_m, ", le combat commence !")
	compteur_tour := 0

	if a.joueur.initiative > a.ennemi.initiative {
		for !(a.joueur.is_dead()) || !(a.ennemi.is_dead()) {
			compteur_tour++
			a.joueur.attaquer(a.info_ennemi)
			if !(a.ennemi.is_dead()) {
				a.ennemi.attaquer(a.info_joueur, compteur_tour)
			} else {
				a.joueur_gagne()
			}
		}
	} else {
		for !(a.joueur.is_dead()) || !(a.ennemi.is_dead()) {

			compteur_tour++
			a.ennemi.attaquer(a.info_joueur, compteur_tour)
			fmt.Println("Le monstre attaque !")
			if !(a.joueur.is_dead()) {
				a.joueur.attaquer(a.info_ennemi)
			} else {
				a.joueur.dead()
				break
			}
		}
	}
}

func crea_perso() Personnage {
	fmt.Println("\nVous allez créer un personnage.")
	fmt.Println("Quel est le nom de votre personnage ?")
	var nom1 string
	fmt.Scan(&nom1)
	var nom string
	if !IsAlpha(nom1) {
		fmt.Println("Le nom de votre personnage ne peut contenir que des lettres.\nQuel est le nom de votre personnage ?")
		fmt.Scan(&nom)
	} else {
		nom = Capitalize(nom1)
	}
	nom = Capitalize(nom)
	fmt.Println("Quel est la classe de votre personnage ? \n 1 : Humain \n 2 : Elfe \n 3 : Nain")
	var classe int
	fmt.Scan(&classe)
	var classe_p string
	var pv_maximum int
	switch classe {
	case 1:
		classe_p = "Humain"
		pv_maximum = 100
	case 2:
		classe_p = "Elfe"
		pv_maximum = 80
	case 3:
		classe_p = "Nain"
		pv_maximum = 120
	}

	var p1 Personnage

	p1.Init(nom, classe_p, 1, pv_maximum, pv_maximum/2, []item{}, []competence{}, equipement{}, 100, 0, 100, 100, 0)
	return p1
}
