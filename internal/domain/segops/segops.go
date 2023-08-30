package segops

import "log"

func UpdatedSegments(oldSeg, toAdd, toRemove []string) *[]string {
	var newSeg []string
	set := toMap(oldSeg)
	for _, s := range toRemove {
		delete(set, s)
	}

	for _, s := range toAdd {
		set[s] = true
	}

	for k := range set {
		newSeg = append(newSeg, k)
	}

	return &newSeg
}

func toMap(s []string) map[string]bool {
	log.Println(s)
	set := make(map[string]bool)

	for _, m := range s {
		set[m] = true
	}

	return set
}
