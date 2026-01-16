-- Drop unused tables from old architecture
-- These tables were used by the layer1 block scanner but are no longer needed

DROP TABLE IF EXISTS mesh_status CASCADE;
DROP TABLE IF EXISTS depth_state CASCADE;
DROP TABLE IF EXISTS mesh_link CASCADE;
DROP TABLE IF EXISTS price_estimate CASCADE;
DROP TABLE IF EXISTS ooconfig CASCADE;
