package aggregates

// Facet Facet-related aggregation stages categorize and group incoming documents.
// Specify any of the following facet-related stages within different $facet
// sub-pipeline's <stage> to perform a multi-faceted aggregation: $bucket
// $bucketAuto $sortByCount Other aggregation stages can also be used with
// $facet with the following exceptions: $collStats $facet $geoNear $indexStats
// $out $merge $planCacheStats Each sub-pipeline within $facet is passed the
// exact same set of input documents. These sub-pipelines are completely
// independent of one another and the document array output by each is stored in
// separate fields in the output document. The output of one sub-pipeline can
// not be used as the input for a different sub-pipeline within the same $facet
// stage. If further aggregations are required, add additional stages after
// $facet and specify the field name, <outputField>, of the desired sub-pipeline
// output.
type Facet struct {
	name     string
	pipeline []Stage
}

func NewFacet(name string, pipeline []Stage) Facet {
	return Facet{
		name:     name,
		pipeline: pipeline,
	}
}
