CREATE TRIGGER erc1155_transf_insert AFTER INSERT ON erc1155_transf FOR EACH ROW EXECUTE PROCEDURE on_erc1155_transf_insert();
CREATE TRIGGER erc1155_transf_delete AFTER DELETE ON erc1155_transf FOR EACH ROW EXECUTE PROCEDURE on_erc1155_transf_delete();
CREATE TRIGGER erc1155_batch_insert AFTER INSERT ON erc1155_batch FOR EACH ROW EXECUTE PROCEDURE on_erc1155_batch_insert();
CREATE TRIGGER erc1155_batch_delete AFTER DELETE ON erc1155_batch FOR EACH ROW EXECUTE PROCEDURE on_erc1155_batch_delete();
CREATE TRIGGER erc1155_bal_insert AFTER INSERT ON erc1155_bal FOR EACH ROW EXECUTE PROCEDURE on_erc1155_bal_insert();
CREATE TRIGGER erc1155_bal_delete AFTER DELETE ON erc1155_bal FOR EACH ROW EXECUTE PROCEDURE on_erc1155_bal_delete();
