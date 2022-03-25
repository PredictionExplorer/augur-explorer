CREATE TRIGGER balv2_swf_hist_insert AFTER INSERT ON swf_hist FOR EACH ROW EXECUTE PROCEDURE on_swf_hist_insert();
CREATE TRIGGER balv2_swf_hist_delete AFTER DELETE ON swf_hist FOR EACH ROW EXECUTE PROCEDURE on_swf_hist_delete();
