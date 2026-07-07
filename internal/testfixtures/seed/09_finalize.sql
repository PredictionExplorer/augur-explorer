-- Advance sequences past explicitly assigned ids so runtime inserts (e.g. the
-- mutation-route tests) never collide with fixture rows.
SELECT setval(pg_get_serial_sequence('address', 'address_id'), (SELECT MAX(address_id) FROM address));
SELECT setval(pg_get_serial_sequence('transaction', 'id'), (SELECT MAX(id) FROM transaction));
SELECT setval(pg_get_serial_sequence('evt_log', 'id'), (SELECT MAX(id) FROM evt_log));
SELECT setval(pg_get_serial_sequence('cg_bid', 'id'), (SELECT MAX(id) FROM cg_bid));
