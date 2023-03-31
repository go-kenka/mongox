package expression

var (
	Now         Variable = Let("NOW")          // since 4.2
	ClusterTime Variable = Let("CLUSTER_TIME") // since 4.2
	Root        Variable = Let("ROOT")
	Current     Variable = Let("CURRENT")
	Remove      Variable = Let("REMOVE")
	Descend     Variable = Let("DESCEND")
	Prune       Variable = Let("PRUNE")
	Keep        Variable = Let("KEEP")
	SearchMeta  Variable = Let("SEARCH_META")
)
