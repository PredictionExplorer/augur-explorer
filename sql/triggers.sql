CREATE TRIGGER update_oi_insert AFTER INSERT ON oi_chg FOR EACH ROW EXECUTE PROCEDURE update_oi_on_insert();
CREATE TRIGGER update_vol_insert AFTER INSERT ON volume FOR EACH ROW EXECUTE PROCEDURE update_vol_on_insert();
CREATE TRIGGER update_mkt_fin_insert AFTER INSERT ON mkt_fin FOR EACH ROW EXECUTE PROCEDURE update_mkt_fin_on_insert();
