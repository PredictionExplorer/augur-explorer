CREATE TRIGGER balv2_swf_hist_insert AFTER INSERT ON swf_hist FOR EACH ROW EXECUTE PROCEDURE on_swf_hist_insert();
CREATE TRIGGER balv2_swf_hist_delete AFTER DELETE ON swf_hist FOR EACH ROW EXECUTE PROCEDURE on_swf_hist_delete();
CREATE TRIGGER poolhist_insert AFTER INSERT ON pool_hist FOR EACH ROW EXECUTE PROCEDURE on_poolhist_insert();
CREATE TRIGGER poolhist_delete AFTER DELETE ON pool_hist FOR EACH ROW EXECUTE PROCEDURE on_poolhist_delete();
CREATE TRIGGER tokbal_insert AFTER INSERT ON tok_bal FOR EACH ROW EXECUTE PROCEDURE on_tokbal_insert();
CREATE TRIGGER tokbal_delete AFTER DELETE ON tok_bal FOR EACH ROW EXECUTE PROCEDURE on_tokbal_delete();
CREATE TRIGGER bpt_transf_insert AFTER INSERT ON bpt_transf FOR EACH ROW EXECUTE PROCEDURE on_bpt_transf_insert();
CREATE TRIGGER bpt_transf_delete AFTER DELETE ON bpt_transf FOR EACH ROW EXECUTE PROCEDURE on_bpt_transf_delete();
