// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package datasource_project

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ProjectDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"description": schema.StringAttribute{
				Computed: true,
			},
			"environments": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"created_at": schema.StringAttribute{
							Computed: true,
						},
						"description": schema.StringAttribute{
							Computed: true,
						},
						"id": schema.Int64Attribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"project_id": schema.Int64Attribute{
							Computed: true,
						},
						"updated_at": schema.StringAttribute{
							Computed: true,
						},
					},
					CustomType: EnvironmentsType{
						ObjectType: types.ObjectType{
							AttrTypes: EnvironmentsValue{}.AttributeTypes(ctx),
						},
					},
				},
				Computed:            true,
				Description:         "The environments of the project.",
				MarkdownDescription: "The environments of the project.",
			},
			"id": schema.Int64Attribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"uuid": schema.StringAttribute{
				Required:            true,
				Description:         "Project UUID",
				MarkdownDescription: "Project UUID",
			},
		},
	}
}

type ProjectModel struct {
	Description  types.String `tfsdk:"description"`
	Environments types.List   `tfsdk:"environments"`
	Id           types.Int64  `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	Uuid         types.String `tfsdk:"uuid"`
}

var _ basetypes.ObjectTypable = EnvironmentsType{}

type EnvironmentsType struct {
	basetypes.ObjectType
}

func (t EnvironmentsType) Equal(o attr.Type) bool {
	other, ok := o.(EnvironmentsType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t EnvironmentsType) String() string {
	return "EnvironmentsType"
}

func (t EnvironmentsType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	createdAtAttribute, ok := attributes["created_at"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`created_at is missing from object`)

		return nil, diags
	}

	createdAtVal, ok := createdAtAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`created_at expected to be basetypes.StringValue, was: %T`, createdAtAttribute))
	}

	descriptionAttribute, ok := attributes["description"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`description is missing from object`)

		return nil, diags
	}

	descriptionVal, ok := descriptionAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`description expected to be basetypes.StringValue, was: %T`, descriptionAttribute))
	}

	idAttribute, ok := attributes["id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`id is missing from object`)

		return nil, diags
	}

	idVal, ok := idAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`id expected to be basetypes.Int64Value, was: %T`, idAttribute))
	}

	nameAttribute, ok := attributes["name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`name is missing from object`)

		return nil, diags
	}

	nameVal, ok := nameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`name expected to be basetypes.StringValue, was: %T`, nameAttribute))
	}

	projectIdAttribute, ok := attributes["project_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`project_id is missing from object`)

		return nil, diags
	}

	projectIdVal, ok := projectIdAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`project_id expected to be basetypes.Int64Value, was: %T`, projectIdAttribute))
	}

	updatedAtAttribute, ok := attributes["updated_at"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`updated_at is missing from object`)

		return nil, diags
	}

	updatedAtVal, ok := updatedAtAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`updated_at expected to be basetypes.StringValue, was: %T`, updatedAtAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return EnvironmentsValue{
		CreatedAt:   createdAtVal,
		Description: descriptionVal,
		Id:          idVal,
		Name:        nameVal,
		ProjectId:   projectIdVal,
		UpdatedAt:   updatedAtVal,
		state:       attr.ValueStateKnown,
	}, diags
}

func NewEnvironmentsValueNull() EnvironmentsValue {
	return EnvironmentsValue{
		state: attr.ValueStateNull,
	}
}

func NewEnvironmentsValueUnknown() EnvironmentsValue {
	return EnvironmentsValue{
		state: attr.ValueStateUnknown,
	}
}

func NewEnvironmentsValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (EnvironmentsValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing EnvironmentsValue Attribute Value",
				"While creating a EnvironmentsValue value, a missing attribute value was detected. "+
					"A EnvironmentsValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("EnvironmentsValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid EnvironmentsValue Attribute Type",
				"While creating a EnvironmentsValue value, an invalid attribute value was detected. "+
					"A EnvironmentsValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("EnvironmentsValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("EnvironmentsValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra EnvironmentsValue Attribute Value",
				"While creating a EnvironmentsValue value, an extra attribute value was detected. "+
					"A EnvironmentsValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra EnvironmentsValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewEnvironmentsValueUnknown(), diags
	}

	createdAtAttribute, ok := attributes["created_at"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`created_at is missing from object`)

		return NewEnvironmentsValueUnknown(), diags
	}

	createdAtVal, ok := createdAtAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`created_at expected to be basetypes.StringValue, was: %T`, createdAtAttribute))
	}

	descriptionAttribute, ok := attributes["description"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`description is missing from object`)

		return NewEnvironmentsValueUnknown(), diags
	}

	descriptionVal, ok := descriptionAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`description expected to be basetypes.StringValue, was: %T`, descriptionAttribute))
	}

	idAttribute, ok := attributes["id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`id is missing from object`)

		return NewEnvironmentsValueUnknown(), diags
	}

	idVal, ok := idAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`id expected to be basetypes.Int64Value, was: %T`, idAttribute))
	}

	nameAttribute, ok := attributes["name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`name is missing from object`)

		return NewEnvironmentsValueUnknown(), diags
	}

	nameVal, ok := nameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`name expected to be basetypes.StringValue, was: %T`, nameAttribute))
	}

	projectIdAttribute, ok := attributes["project_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`project_id is missing from object`)

		return NewEnvironmentsValueUnknown(), diags
	}

	projectIdVal, ok := projectIdAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`project_id expected to be basetypes.Int64Value, was: %T`, projectIdAttribute))
	}

	updatedAtAttribute, ok := attributes["updated_at"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`updated_at is missing from object`)

		return NewEnvironmentsValueUnknown(), diags
	}

	updatedAtVal, ok := updatedAtAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`updated_at expected to be basetypes.StringValue, was: %T`, updatedAtAttribute))
	}

	if diags.HasError() {
		return NewEnvironmentsValueUnknown(), diags
	}

	return EnvironmentsValue{
		CreatedAt:   createdAtVal,
		Description: descriptionVal,
		Id:          idVal,
		Name:        nameVal,
		ProjectId:   projectIdVal,
		UpdatedAt:   updatedAtVal,
		state:       attr.ValueStateKnown,
	}, diags
}

func NewEnvironmentsValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) EnvironmentsValue {
	object, diags := NewEnvironmentsValue(attributeTypes, attributes)

	if diags.HasError() {
		// This could potentially be added to the diag package.
		diagsStrings := make([]string, 0, len(diags))

		for _, diagnostic := range diags {
			diagsStrings = append(diagsStrings, fmt.Sprintf(
				"%s | %s | %s",
				diagnostic.Severity(),
				diagnostic.Summary(),
				diagnostic.Detail()))
		}

		panic("NewEnvironmentsValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t EnvironmentsType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewEnvironmentsValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewEnvironmentsValueUnknown(), nil
	}

	if in.IsNull() {
		return NewEnvironmentsValueNull(), nil
	}

	attributes := map[string]attr.Value{}

	val := map[string]tftypes.Value{}

	err := in.As(&val)

	if err != nil {
		return nil, err
	}

	for k, v := range val {
		a, err := t.AttrTypes[k].ValueFromTerraform(ctx, v)

		if err != nil {
			return nil, err
		}

		attributes[k] = a
	}

	return NewEnvironmentsValueMust(EnvironmentsValue{}.AttributeTypes(ctx), attributes), nil
}

func (t EnvironmentsType) ValueType(ctx context.Context) attr.Value {
	return EnvironmentsValue{}
}

var _ basetypes.ObjectValuable = EnvironmentsValue{}

type EnvironmentsValue struct {
	CreatedAt   basetypes.StringValue `tfsdk:"created_at"`
	Description basetypes.StringValue `tfsdk:"description"`
	Id          basetypes.Int64Value  `tfsdk:"id"`
	Name        basetypes.StringValue `tfsdk:"name"`
	ProjectId   basetypes.Int64Value  `tfsdk:"project_id"`
	UpdatedAt   basetypes.StringValue `tfsdk:"updated_at"`
	state       attr.ValueState
}

func (v EnvironmentsValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 6)

	var val tftypes.Value
	var err error

	attrTypes["created_at"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["description"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["id"] = basetypes.Int64Type{}.TerraformType(ctx)
	attrTypes["name"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["project_id"] = basetypes.Int64Type{}.TerraformType(ctx)
	attrTypes["updated_at"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 6)

		val, err = v.CreatedAt.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["created_at"] = val

		val, err = v.Description.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["description"] = val

		val, err = v.Id.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["id"] = val

		val, err = v.Name.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["name"] = val

		val, err = v.ProjectId.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["project_id"] = val

		val, err = v.UpdatedAt.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["updated_at"] = val

		if err := tftypes.ValidateValue(objectType, vals); err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		return tftypes.NewValue(objectType, vals), nil
	case attr.ValueStateNull:
		return tftypes.NewValue(objectType, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(objectType, tftypes.UnknownValue), nil
	default:
		panic(fmt.Sprintf("unhandled Object state in ToTerraformValue: %s", v.state))
	}
}

func (v EnvironmentsValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v EnvironmentsValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v EnvironmentsValue) String() string {
	return "EnvironmentsValue"
}

func (v EnvironmentsValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributeTypes := map[string]attr.Type{
		"created_at":  basetypes.StringType{},
		"description": basetypes.StringType{},
		"id":          basetypes.Int64Type{},
		"name":        basetypes.StringType{},
		"project_id":  basetypes.Int64Type{},
		"updated_at":  basetypes.StringType{},
	}

	if v.IsNull() {
		return types.ObjectNull(attributeTypes), diags
	}

	if v.IsUnknown() {
		return types.ObjectUnknown(attributeTypes), diags
	}

	objVal, diags := types.ObjectValue(
		attributeTypes,
		map[string]attr.Value{
			"created_at":  v.CreatedAt,
			"description": v.Description,
			"id":          v.Id,
			"name":        v.Name,
			"project_id":  v.ProjectId,
			"updated_at":  v.UpdatedAt,
		})

	return objVal, diags
}

func (v EnvironmentsValue) Equal(o attr.Value) bool {
	other, ok := o.(EnvironmentsValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.CreatedAt.Equal(other.CreatedAt) {
		return false
	}

	if !v.Description.Equal(other.Description) {
		return false
	}

	if !v.Id.Equal(other.Id) {
		return false
	}

	if !v.Name.Equal(other.Name) {
		return false
	}

	if !v.ProjectId.Equal(other.ProjectId) {
		return false
	}

	if !v.UpdatedAt.Equal(other.UpdatedAt) {
		return false
	}

	return true
}

func (v EnvironmentsValue) Type(ctx context.Context) attr.Type {
	return EnvironmentsType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v EnvironmentsValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"created_at":  basetypes.StringType{},
		"description": basetypes.StringType{},
		"id":          basetypes.Int64Type{},
		"name":        basetypes.StringType{},
		"project_id":  basetypes.Int64Type{},
		"updated_at":  basetypes.StringType{},
	}
}
