package aggregates

import (
	"github.com/go-kenka/mongox/examples/data/bsonx"
	"github.com/go-kenka/mongox/examples/data/model/densify"
	"github.com/go-kenka/mongox/examples/data/model/expressions"
	"github.com/go-kenka/mongox/examples/data/model/fill"
	"github.com/go-kenka/mongox/examples/data/model/geojson"
	"github.com/go-kenka/mongox/examples/data/model/search"
	"github.com/go-kenka/mongox/examples/data/model/window"
)

func AddFields(fields ...Field[expressions.TExpression]) bsonx.Bson {
	return NewFieldsStage("$addFields", fields)
}
func Set(fields ...Field[expressions.TExpression]) bsonx.Bson {
	return NewFieldsStage("$set", fields)
}
func UnSet(fields ...string) bsonx.Bson {
	if len(fields) == 1 {
		return bsonx.NewBsonDocument("$unset", bsonx.NewBsonString(fields[0]))
	}

	array := bsonx.NewBsonArray()
	for _, field := range fields {
		array.Append(bsonx.NewBsonString(field))
	}

	return bsonx.NewBsonDocument("$unset", array)
}
func Bucket(groupBy expressions.TExpression, boundaries []expressions.TBoundary, options AggBucketOptions) bsonx.Bson {
	return NewBucketStage(groupBy, boundaries, options)
}
func BucketAuto(groupBy expressions.TExpression, buckets int32, options BucketAutoOptions) bsonx.Bson {
	return NewBucketAutoStage(groupBy, buckets, options)
}
func Count(field string) bsonx.Bson {
	return bsonx.NewBsonDocument("$count", bsonx.NewBsonString(field))
}
func Match(filter bsonx.Bson) bsonx.Bson {
	return NewSimplePipelineStage("$match", filter)
}
func Project(projection bsonx.Bson) bsonx.Bson {
	return NewSimplePipelineStage("$project", projection)
}
func Sort(sort bsonx.Bson) bsonx.Bson {
	return NewSimplePipelineStage("$sort", sort)
}
func SortByCount(filter expressions.TExpression) bsonx.Bson {
	return NewSortByCountStage(filter)
}
func Skip(skip int32) bsonx.Bson {
	return bsonx.NewBsonDocument("$skip", bsonx.NewBsonInt32(skip))
}
func Limit(limit int32) bsonx.Bson {
	return bsonx.NewBsonDocument("$limit", bsonx.NewBsonInt32(limit))
}
func Lookup(from, localField, foreignField, as string) bsonx.Bson {
	return bsonx.NewBsonDocument("$lookup", bsonx.NewBsonDocument("from", bsonx.NewBsonString(from)).
		Append("localField", bsonx.NewBsonString(localField)).
		Append("foreignField", bsonx.NewBsonString(foreignField)).
		Append("as", bsonx.NewBsonString(as)))
}
func LookupWithPipe(from, as string, let []Variable[expressions.TExpression], pipeline []bsonx.Bson) bsonx.Bson {
	return NewLookupStage(from, let, pipeline, as)
}
func Facets(facets ...Facet) bsonx.Bson {
	return NewFacetStage(facets)
}
func GraphLookup(
	from string,
	startWith expressions.TExpression,
	connectFromField string,
	connectToField string,
	as string,
	options GraphLookupOptions,
) bsonx.Bson {
	return NewGraphLookupStage(from, startWith, connectFromField, connectToField, as, options)
}
func Group(id expressions.TExpression, fieldAccumulators ...bsonx.BsonField) bsonx.Bson {
	return NewGroupStage(id, fieldAccumulators)
}
func UnionWith(collection string, pipeline ...bsonx.Bson) bsonx.Bson {
	return NewUnionWithStage(collection, pipeline...)
}
func Unwind(fieldName string, unwindOptions *UnwindOptions) bsonx.Bson {
	if unwindOptions == nil {
		return bsonx.NewBsonDocument("$unwind", bsonx.NewBsonString(fieldName))
	}
	options := bsonx.NewBsonDocument("path", bsonx.NewBsonString(fieldName))
	if unwindOptions.HasPreserveNullAndEmptyArrays() {
		options.Append("preserveNullAndEmptyArrays", bsonx.NewBsonBoolean(unwindOptions.PreserveNullAndEmptyArrays()))
	}
	if unwindOptions.HasIncludeArrayIndex() {
		options.Append("includeArrayIndex", bsonx.NewBsonString(unwindOptions.IncludeArrayIndex()))
	}
	return bsonx.NewBsonDocument("$unwind", options)
}
func Out(databaseName, collectionName string) bsonx.Bson {
	if len(databaseName) == 0 {
		return bsonx.NewBsonDocument("$out", bsonx.NewBsonString(collectionName))
	}
	return bsonx.NewBsonDocument("$out", bsonx.NewBsonDocument("db", bsonx.NewBsonString(databaseName)).
		Append("coll", bsonx.NewBsonString(collectionName)))
}
func OutWithBson(destination bsonx.Bson) bsonx.Bson {
	return NewSimplePipelineStage("$out", destination)
}
func Merge(collectionName string, options MergeOptions) bsonx.Bson {
	return NewMergeStage(bsonx.NewBsonString(collectionName), options)
}
func MergeWithNameSpace(namespace MongoNamespace, options MergeOptions) bsonx.Bson {
	return NewMergeStage(bsonx.NewBsonDocument("db", bsonx.NewBsonString(namespace.databaseName)).
		Append("coll", bsonx.NewBsonString(namespace.collectionName)), options)
}
func ReplaceRoot(value expressions.TExpression) bsonx.Bson {
	return NewReplaceStage(value, false)
}
func ReplaceWith(value expressions.TExpression) bsonx.Bson {
	return NewReplaceStage(value, true)
}
func Sample(size int32) bsonx.Bson {
	return bsonx.NewBsonDocument("$sample", bsonx.NewBsonDocument("size", bsonx.NewBsonInt32(size)))
}
func SetWindowFields(partitionBy expressions.TExpression, sortBy bsonx.Bson, output []window.WindowOutputField) bsonx.Bson {
	return NewSetWindowFieldsStage(partitionBy, sortBy, output)
}
func Densify(field string, dRange densify.DensifyRange) bsonx.Bson {
	return NewDensifyStage(field, dRange, densify.DefaultDensifyOptions)
}
func Fill(options fill.FillOptions, output []fill.FillOutputField) bsonx.Bson {
	return NewFillStage(options, output)
}
func Search(operator search.SearchOperator, options search.SearchOptions) bsonx.Bson {
	return NewSearchStage("$search", operator, options)
}
func SearchWithCollector(collector search.SearchCollector, options search.SearchOptions) bsonx.Bson {
	return NewSearchStage("$search", collector, options)
}
func SearchMeta(operator search.SearchOperator, options search.SearchOptions) bsonx.Bson {
	return NewSearchStage("$searchMeta", operator, options)
}
func SearchMetaWithCollector(collector search.SearchCollector, options search.SearchOptions) bsonx.Bson {
	return NewSearchStage("$searchMeta", collector, options)
}
func GeoNear(near geojson.Point, distanceField string, options geojson.GeoNearOptions) bsonx.Bson {
	return NewGeoNearStage(near, distanceField, options)
}
func Documents(documents []bsonx.Bson) bsonx.Bson {
	return NewDocumentsStage(documents)
}

type SimplePipelineStage struct {
	name  string
	value bsonx.Bson
}

func NewSimplePipelineStage(name string, value bsonx.Bson) SimplePipelineStage {
	return SimplePipelineStage{
		name:  name,
		value: value,
	}
}

func (s SimplePipelineStage) ToBsonDocument() bsonx.BsonDocument {
	return bsonx.NewBsonDocument(s.name, s.value.ToBsonDocument())
}

func (s SimplePipelineStage) Document() bsonx.Document {
	return s.ToBsonDocument().Document()
}

type BucketStage[T expressions.TExpression, B expressions.TBoundary] struct {
	groupBy    T
	boundaries []B
	options    AggBucketOptions
}

func NewBucketStage[T expressions.TExpression, B expressions.TBoundary](groupBy T, boundaries []B, options AggBucketOptions) BucketStage[T, B] {
	return BucketStage[T, B]{
		groupBy:    groupBy,
		boundaries: boundaries,
		options:    options,
	}
}

func (bs BucketStage[T, B]) ToBsonDocument() bsonx.BsonDocument {
	b := bsonx.NewEmptyDoc()
	data := bsonx.NewBsonDocument("groupBy", bs.groupBy)

	var boundaries bsonx.BsonArray
	for _, boundary := range bs.boundaries {
		boundaries.Append(boundary)
	}

	data.Append("boundaries", boundaries)

	defaultBucket := bs.options.GetDefaultBucket()
	if defaultBucket != "" {
		data.Append("default", bsonx.NewBsonString(defaultBucket))
	}

	output := bs.options.GetOutPut()

	if len(output) > 0 {
		out := bsonx.NewEmptyDoc()
		for _, field := range output {
			out.Append(field.GetName(), field.GetValue().ToBsonDocument())
		}
		data.Append("output", out)
	}

	b.Append("$bucket", data)
	return b
}
func (bs BucketStage[T, B]) Document() bsonx.Document {
	return bs.ToBsonDocument().Document()
}

type BucketAutoStage[T expressions.TExpression] struct {
	groupBy T
	buckets int32
	options BucketAutoOptions
}

func NewBucketAutoStage[T expressions.TExpression](groupBy T, buckets int32, options BucketAutoOptions) BucketAutoStage[T] {
	return BucketAutoStage[T]{
		groupBy: groupBy,
		buckets: buckets,
		options: options,
	}
}

func (bs BucketAutoStage[T]) ToBsonDocument() bsonx.BsonDocument {
	b := bsonx.NewEmptyDoc()
	data := bsonx.NewBsonDocument("groupBy", bs.groupBy)

	data.Append("buckets", bsonx.NewBsonInt32(bs.buckets))

	output := bs.options.GetOutPut()

	if len(output) > 0 {
		out := bsonx.NewEmptyDoc()
		for _, field := range output {
			out.Append(field.GetName(), field.GetValue().ToBsonDocument())
		}
		data.Append("output", out)
	}

	granularity := bs.options.GetGranularity()
	if len(granularity) > 0 {
		data.Append("granularity", bsonx.NewBsonString(granularity))
	}

	b.Append("$bucketAuto", data)
	return b
}
func (bs BucketAutoStage[T]) Document() bsonx.Document {
	return bs.ToBsonDocument().Document()
}

type LookupStage[T expressions.TExpression] struct {
	from     string
	let      []Variable[T]
	pipeline []bsonx.Bson
	as       string
}

func NewLookupStage[T expressions.TExpression](
	from string,
	let []Variable[T],
	pipeline []bsonx.Bson,
	as string,
) LookupStage[T] {
	return LookupStage[T]{
		from:     from,
		let:      let,
		pipeline: pipeline,
		as:       as,
	}
}

func (bs LookupStage[T]) ToBsonDocument() bsonx.BsonDocument {
	b := bsonx.NewEmptyDoc()
	data := bsonx.NewBsonDocument("form", bsonx.NewBsonString(bs.from))

	if len(bs.let) > 0 {
		let := bsonx.NewEmptyDoc()
		for _, field := range bs.let {
			let.Append(field.GetName(), field.GetValue())
		}
		data.Append("let", let)
	}

	if len(bs.pipeline) > 0 {
		var pipeline bsonx.BsonArray
		for _, p := range bs.pipeline {
			pipeline.Append(p.ToBsonDocument())
		}
		data.Append("pipeline", pipeline)
	}
	data.Append("as", bsonx.NewBsonString(bs.as))
	b.Append("$lookup", data)
	return b
}
func (bs LookupStage[T]) Document() bsonx.Document {
	return bs.ToBsonDocument().Document()
}

type GraphLookupStage[T expressions.TExpression] struct {
	from             string
	startWith        T
	connectFromField string
	connectToField   string
	as               string
	options          GraphLookupOptions
}

func NewGraphLookupStage[T expressions.TExpression](
	from string,
	startWith T,
	connectFromField string,
	connectToField string,
	as string,
	options GraphLookupOptions,
) GraphLookupStage[T] {
	return GraphLookupStage[T]{
		from:             from,
		startWith:        startWith,
		connectFromField: connectFromField,
		connectToField:   connectToField,
		as:               as,
		options:          options,
	}
}

func (bs GraphLookupStage[T]) ToBsonDocument() bsonx.BsonDocument {
	b := bsonx.NewEmptyDoc()
	data := bsonx.NewBsonDocument("form", bsonx.NewBsonString(bs.from))

	data.Append("startWith", bs.startWith)
	data.Append("connectFromField", bsonx.NewBsonString(bs.connectFromField))
	data.Append("connectToField", bsonx.NewBsonString(bs.connectToField))
	data.Append("as", bsonx.NewBsonString(bs.as))

	if bs.options.GetMaxDepth() > 0 {
		data.Append("maxDepth", bsonx.NewBsonInt32(bs.options.GetMaxDepth()))
	}
	if bs.options.GetDepthField() != "" {
		data.Append("depthField", bsonx.NewBsonString(bs.options.GetDepthField()))
	}
	if bs.options.GetRestrictSearchWithMatch() != nil {
		data.Append("restrictSearchWithMatch", bs.options.GetRestrictSearchWithMatch().ToBsonDocument())
	}
	b.Append("$graphLookup", data)
	return b
}

func (bs GraphLookupStage[T]) Document() bsonx.Document {
	return bs.ToBsonDocument().Document()
}

type GroupStage[T expressions.TExpression] struct {
	id                T
	fieldAccumulators []bsonx.BsonField
}

func NewGroupStage[T expressions.TExpression](id T, fieldAccumulators []bsonx.BsonField) GroupStage[T] {
	return GroupStage[T]{
		id:                id,
		fieldAccumulators: fieldAccumulators,
	}
}

func (bs GroupStage[T]) ToBsonDocument() bsonx.BsonDocument {
	b := bsonx.NewEmptyDoc()
	data := bsonx.NewBsonDocument("_id", bs.id)

	if len(bs.fieldAccumulators) > 0 {
		for _, field := range bs.fieldAccumulators {
			data.Append(field.GetName(), field.GetValue().ToBsonDocument())
		}
	}

	b.Append("$group", data)
	return b
}

func (bs GroupStage[T]) Document() bsonx.Document {
	return bs.ToBsonDocument().Document()
}

type SortByCountStage[T expressions.TExpression] struct {
	filter T
}

func NewSortByCountStage[T expressions.TExpression](filter T) SortByCountStage[T] {
	return SortByCountStage[T]{
		filter: filter,
	}
}

func (bs SortByCountStage[T]) ToBsonDocument() bsonx.BsonDocument {
	return bsonx.NewBsonDocument("$sortByCount", bs.filter)
}
func (bs SortByCountStage[T]) Document() bsonx.Document {
	return bs.ToBsonDocument().Document()
}

type FacetStage struct {
	facets []Facet
}

func NewFacetStage(facets []Facet) FacetStage {
	return FacetStage{
		facets: facets,
	}
}

func (bs FacetStage) ToBsonDocument() bsonx.BsonDocument {
	b := bsonx.NewEmptyDoc()
	data := bsonx.NewEmptyDoc()

	if len(bs.facets) > 0 {
		for _, f := range bs.facets {
			var pipeline bsonx.BsonArray
			for _, p := range f.pipeline {
				pipeline.Append(p.ToBsonDocument())
			}

			data.Append(f.name, pipeline)
		}
	}
	b.Append("$facet", data)
	return b
}

func (bs FacetStage) Document() bsonx.Document {
	return bs.ToBsonDocument().Document()
}

type FieldsStage[T expressions.TExpression] struct {
	fields    []Field[T]
	stageName string
}

func NewFieldsStage[T expressions.TExpression](stageName string, fields []Field[T]) FieldsStage[T] {
	return FieldsStage[T]{
		fields:    fields,
		stageName: stageName,
	}
}

func (f FieldsStage[T]) ToBsonDocument() bsonx.BsonDocument {
	b := bsonx.NewEmptyDoc()
	data := bsonx.NewEmptyDoc()
	for _, field := range f.fields {
		data.Append(field.name, field.value)
	}
	b.Append(f.stageName, data)
	return b
}

func (f FieldsStage[T]) Document() bsonx.Document {
	return f.ToBsonDocument().Document()
}

type ReplaceStage[T expressions.TExpression] struct {
	value       T
	replaceWith bool
}

func NewReplaceStage[T expressions.TExpression](value T, replaceWith bool) ReplaceStage[T] {
	return ReplaceStage[T]{
		value:       value,
		replaceWith: replaceWith,
	}
}

func (f ReplaceStage[T]) ToBsonDocument() bsonx.BsonDocument {
	if f.replaceWith {
		return bsonx.NewBsonDocument("$replaceWith", f.value)
	}
	b := bsonx.NewEmptyDoc()
	data := bsonx.NewBsonDocument("newRoot", f.value)
	b.Append("$replaceRoot", data)
	return b
}

func (f ReplaceStage[T]) Document() bsonx.Document {
	return f.ToBsonDocument().Document()
}

type MergeStage struct {
	intoValue bsonx.IBsonValue
	options   MergeOptions
}

func NewMergeStage(intoValue bsonx.IBsonValue, options MergeOptions) MergeStage {
	return MergeStage{
		intoValue: intoValue,
		options:   options,
	}
}

func (f MergeStage) ToBsonDocument() bsonx.BsonDocument {
	b := bsonx.NewEmptyDoc()
	data := bsonx.NewEmptyDoc()
	if f.intoValue.IsString() {
		data.Append("into", f.intoValue)
	} else {
		into := bsonx.NewEmptyDoc()
		into.Append("db", f.intoValue.AsDocument().GetString("db"))
		into.Append("coll", f.intoValue.AsDocument().GetString("coll"))
		data.Append("into", into)
	}

	if len(f.options.uniqueIdentifier) > 0 {
		if len(f.options.uniqueIdentifier) == 1 {
			data.Append("on", bsonx.NewBsonString(f.options.uniqueIdentifier[0]))
		} else {
			var uniqueIdentifier bsonx.BsonArray
			for _, s := range f.options.uniqueIdentifier {
				uniqueIdentifier.Append(bsonx.NewBsonString(s))
			}
			data.Append("on", uniqueIdentifier)
		}
	}
	if len(f.options.variables) > 0 {
		variables := bsonx.NewEmptyDoc()
		for _, s := range f.options.variables {
			variables.Append(s.GetName(), s.GetValue())
		}
		data.Append("let", variables)
	}

	if f.options.whenMatched != WhenMatchedInvalid {
		switch f.options.whenMatched {
		case WhenMatchedReplace, WhenMatchedKeepExisting, WhenMatchedMerge, WhenMatchedFail:
			data.Append("whenMatched", bsonx.NewBsonString(WhenMatcheds[f.options.whenMatched]))
		case WhenMatchedPipeline:
			pipe := bsonx.NewBsonArray()
			for _, m := range f.options.whenMatchedPipeline {
				pipe.Append(bsonx.NewBsonBoolean(m))
			}
			data.Append("whenMatched", pipe)
		}
	}
	if f.options.whenNotMatched != WhenNotMatchedInvalid {
		data.Append("whenNotMatched", bsonx.NewBsonString(WhenNotMatcheds[f.options.whenNotMatched]))
	}

	b.Append("$merge", data)
	return b
}

func (f MergeStage) Document() bsonx.Document {
	return f.ToBsonDocument().Document()
}

type UnionWithStage struct {
	collection string
	pipeline   []bsonx.Bson
}

func NewUnionWithStage(collection string, pipeline ...bsonx.Bson) UnionWithStage {
	return UnionWithStage{
		collection: collection,
		pipeline:   pipeline,
	}
}

func (f UnionWithStage) ToBsonDocument() bsonx.BsonDocument {
	b := bsonx.NewEmptyDoc()
	data := bsonx.NewEmptyDoc()
	data.Append("coll", bsonx.NewBsonString(f.collection))

	var pipeline bsonx.BsonArray
	for _, s := range f.pipeline {
		pipeline.Append(s.ToBsonDocument())
	}
	data.Append("pipeline", pipeline)

	b.Append("$unionWith", data)
	return b
}

func (f UnionWithStage) Document() bsonx.Document {
	return f.ToBsonDocument().Document()
}

type SetWindowFieldsStage[T expressions.TExpression] struct {
	partitionBy T
	sortBy      bsonx.Bson
	output      []window.WindowOutputField
}

func NewSetWindowFieldsStage[T expressions.TExpression](
	partitionBy T,
	sortBy bsonx.Bson,
	output []window.WindowOutputField,
) SetWindowFieldsStage[T] {
	return SetWindowFieldsStage[T]{
		partitionBy: partitionBy,
		sortBy:      sortBy,
		output:      output,
	}
}

func (f SetWindowFieldsStage[T]) ToBsonDocument() bsonx.BsonDocument {
	b := bsonx.NewEmptyDoc()
	data := bsonx.NewEmptyDoc()
	if f.partitionBy != nil {
		data.Append("partitionBy", f.partitionBy)
	}
	if f.sortBy != nil {
		data.Append("sortBy", f.sortBy.ToBsonDocument())
	}
	output := bsonx.NewEmptyDoc()
	for _, s := range f.output {
		field := s.ToBsonField()
		output.Append(field.GetName(), field.GetValue().ToBsonDocument())
	}
	data.Append("output", output)
	b.Append("$setWindowFields", data)
	return b
}

func (f SetWindowFieldsStage[T]) Document() bsonx.Document {
	return f.ToBsonDocument().Document()
}

// SearchStage TODO: SearchStage未完成
type SearchStage struct {
	name                string
	operatorOrCollector bsonx.Bson
	options             search.SearchOptions
}

func NewSearchStage(
	name string,
	operatorOrCollector bsonx.Bson,
	options search.SearchOptions,
) SearchStage {
	return SearchStage{
		name:                name,
		operatorOrCollector: operatorOrCollector,
		options:             options,
	}
}

func (f SearchStage) ToBsonDocument() bsonx.BsonDocument {
	return bsonx.NewEmptyDoc()
}

func (f SearchStage) Document() bsonx.Document {
	return f.ToBsonDocument().Document()
}

type DensifyStage struct {
	field        string
	densifyRange densify.DensifyRange
	options      densify.DensifyOptions
}

func NewDensifyStage(
	field string,
	densifyRange densify.DensifyRange,
	options densify.DensifyOptions,
) DensifyStage {
	return DensifyStage{
		field:        field,
		densifyRange: densifyRange,
		options:      options,
	}
}

func (f DensifyStage) ToBsonDocument() bsonx.BsonDocument {
	doc := bsonx.NewBsonDocument("field", bsonx.NewBsonString(f.field))
	doc.Append("range", f.densifyRange.ToBsonDocument())
	return bsonx.NewMerged(doc, f.options.ToBsonDocument())
}

func (f DensifyStage) Document() bsonx.Document {
	return f.ToBsonDocument().Document()
}

type FillStage struct {
	output  []fill.FillOutputField
	options fill.FillOptions
}

func NewFillStage(
	options fill.FillOptions,
	output []fill.FillOutputField,
) FillStage {
	return FillStage{
		output:  output,
		options: options,
	}
}

func (f FillStage) ToBsonDocument() bsonx.BsonDocument {
	doc := bsonx.NewEmptyDoc()
	doc = bsonx.NewMerged(doc, f.options.ToBsonDocument())
	outputDoc := bsonx.NewEmptyDoc()
	for _, computation := range f.output {
		computationDoc := computation.ToBsonDocument()
		if computationDoc.Size() == 1 {
			outputDoc = bsonx.NewMerged(outputDoc, computationDoc)
		}
	}

	doc.Append("output", outputDoc)
	return bsonx.NewBsonDocument("$fill", doc)
}
func (f FillStage) Document() bsonx.Document {
	return f.ToBsonDocument().Document()
}

type GeoNearStage struct {
	near          geojson.Point
	distanceField string
	options       geojson.GeoNearOptions
}

func NewGeoNearStage(
	near geojson.Point,
	distanceField string,
	options geojson.GeoNearOptions,
) GeoNearStage {
	return GeoNearStage{
		near:          near,
		distanceField: distanceField,
		options:       options,
	}
}

func (f GeoNearStage) ToBsonDocument() bsonx.BsonDocument {
	doc := bsonx.NewEmptyDoc()

	geoNear := bsonx.NewEmptyDoc()
	geoNear.Append("near", f.near.Encode())
	geoNear.Append("distanceField", bsonx.NewBsonString(f.distanceField))

	for key, value := range f.options.ToBsonDocument().Data() {
		geoNear.Append(key, value)
	}

	doc.Append("$geoNear", geoNear)
	return doc
}

func (f GeoNearStage) Document() bsonx.Document {
	return f.ToBsonDocument().Document()
}

type DocumentsStage struct {
	documents []bsonx.Bson
}

func NewDocumentsStage(documents []bsonx.Bson) DocumentsStage {
	return DocumentsStage{
		documents: documents,
	}
}

func (f DocumentsStage) ToBsonDocument() bsonx.BsonDocument {
	doc := bsonx.NewEmptyDoc()
	documents := bsonx.NewBsonArray()
	for _, value := range f.documents {
		documents.Append(value.ToBsonDocument())
	}

	doc.Append("$documents", documents)
	return doc
}

func (f DocumentsStage) Document() bsonx.Document {
	return f.ToBsonDocument().Document()
}
