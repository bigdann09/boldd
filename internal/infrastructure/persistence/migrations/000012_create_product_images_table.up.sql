CREATE TABLE IF NOT EXISTS product_images (
    product_id UUID REFERENCES products(id) ON DELETE CASCADE,
    product_variant_id UUID REFERENCES product_variants(id) ON DELETE CASCADE,
    image_url TEXT NOT NULL
);