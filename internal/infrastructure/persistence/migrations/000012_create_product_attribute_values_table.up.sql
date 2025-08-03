CREATE TABLE IF NOT EXISTS product_attribute_values (
    product_variant_id UUID REFERENCES product_variants(id) ON DELETE CASCADE,
    attribute_id UUID REFERENCES attributes(id) ON DELETE CASCADE,
    value VARCHAR(20) NOT NULL
);