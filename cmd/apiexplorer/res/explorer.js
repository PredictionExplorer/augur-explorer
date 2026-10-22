// API explorer shell behaviour: sidebar params, hover-card pinning and
// placement, copy buttons, full-hex toggle, quick-jump focus, JSON drawer.
// Plain DOM, no dependencies; every feature degrades to static links.
(function () {
	"use strict";

	var body = document.body;
	var store = window.localStorage;

	function get(k) { try { return store.getItem(k); } catch (e) { return null; } }
	function set(k, v) { try { store.setItem(k, v); } catch (e) { /* private mode */ } }

	// ---- sidebar -------------------------------------------------------
	var navToggle = document.getElementById("nav-toggle");
	if (get("nav-hidden") === "1") body.classList.add("nav-hidden");
	if (navToggle) {
		navToggle.addEventListener("click", function () {
			body.classList.toggle("nav-hidden");
			set("nav-hidden", body.classList.contains("nav-hidden") ? "1" : "0");
		});
	}

	// Param groups: links carry "{name}" (html/template URL-escapes braces
	// to %7Bname%7D in paths and %7bname%7d in query strings); substitute
	// the input value and remember it.
	function applyParam(input) {
		var name = input.getAttribute("data-param");
		var val = input.value.trim();
		var group = input.closest(".nav-group");
		if (!group) return;
		var links = group.querySelectorAll("a");
		for (var i = 0; i < links.length; i++) {
			var a = links[i];
			var tpl = a.getAttribute("data-tpl") || a.getAttribute("href");
			a.setAttribute("data-tpl", tpl);
			var href = tpl.split("%7B" + name + "%7D").join("{" + name + "}")
				.split("%7b" + name + "%7d").join("{" + name + "}");
			if (val) {
				a.setAttribute("href", href.split("{" + name + "}").join(encodeURIComponent(val)));
				a.classList.remove("needs-param");
			} else {
				a.setAttribute("href", "#");
				a.classList.add("needs-param");
			}
		}
	}
	var params = document.querySelectorAll(".nav-param");
	for (var p = 0; p < params.length; p++) {
		(function (input) {
			var key = "param:" + input.getAttribute("data-param");
			var saved = get(key);
			if (saved) input.value = saved;
			applyParam(input);
			input.addEventListener("input", function () {
				set(key, input.value.trim());
				applyParam(input);
			});
			input.addEventListener("keydown", function (e) {
				if (e.key === "Enter") {
					var first = input.closest(".nav-group").querySelector("a:not(.needs-param)");
					if (first) window.location.href = first.getAttribute("href");
				}
			});
		})(params[p]);
	}
	document.addEventListener("click", function (e) {
		var a = e.target.closest("a.needs-param");
		if (!a) return;
		e.preventDefault();
		var input = a.closest(".nav-group").querySelector(".nav-param");
		if (input) input.focus();
	});

	// Prefill sidebar params from the current page's path values when the
	// user has not typed anything yet (e.g. opening a round page).
	(function prefill() {
		var m = window.location.pathname.match(/0x[0-9a-fA-F]{40}/);
		var addrInput = document.querySelector('.nav-param[data-param="user_addr"]');
		if (m && addrInput && !addrInput.value) { addrInput.value = m[0]; applyParam(addrInput); }
		var r = window.location.pathname.match(/\/(?:round\/info|rounds\/info|by_round(?:\/\w+)?|claims\/detail|eth\/round|list\/by_round)\/(\d+)/)
			|| window.location.search.match(/[?&]round=(\d+)/);
		var roundInput = document.querySelector('.nav-param[data-param="round_num"]');
		if (r && roundInput && !roundInput.value) { roundInput.value = r[1]; applyParam(roundInput); }
	})();

	// ---- hover cards ---------------------------------------------------
	function placeCard(hx) {
		var card = hx.querySelector(":scope > .hx-card");
		if (!card) return;
		hx.classList.remove("flip", "up");
		var r = card.getBoundingClientRect();
		if (r.right > window.innerWidth - 8) hx.classList.add("flip");
		if (r.bottom > window.innerHeight - 24 && r.top > r.height) hx.classList.add("up");
	}
	document.addEventListener("mouseover", function (e) {
		var hx = e.target.closest(".hx");
		if (hx) placeCard(hx);
	});
	document.addEventListener("click", function (e) {
		var copy = e.target.closest(".hx-copy");
		if (copy) {
			e.preventDefault();
			e.stopPropagation();
			var text = copy.getAttribute("data-copy") || "";
			var done = function () {
				copy.classList.add("done");
				setTimeout(function () { copy.classList.remove("done"); }, 900);
			};
			if (navigator.clipboard && navigator.clipboard.writeText) {
				navigator.clipboard.writeText(text).then(done, done);
			} else {
				var ta = document.createElement("textarea");
				ta.value = text; document.body.appendChild(ta); ta.select();
				try { document.execCommand("copy"); } catch (err) { /* ignore */ }
				document.body.removeChild(ta); done();
			}
			return;
		}
		var dots = e.target.closest(".hx-dots");
		if (dots) {
			var hx = dots.closest(".hx");
			var pinned = hx.classList.toggle("pin");
			if (pinned) placeCard(hx);
			e.stopPropagation();
			return;
		}
		// Click on a card keeps it; click elsewhere unpins everything.
		if (e.target.closest(".hx-card")) return;
		var pins = document.querySelectorAll(".hx.pin");
		for (var i = 0; i < pins.length; i++) pins[i].classList.remove("pin");
	});

	// ---- full hex toggle -----------------------------------------------
	var hexToggle = document.getElementById("hex-toggle");
	function syncHex() {
		var on = body.classList.contains("full-hex");
		if (hexToggle) hexToggle.classList.toggle("on", on);
	}
	if (get("full-hex") === "1") body.classList.add("full-hex");
	syncHex();
	if (hexToggle) {
		hexToggle.addEventListener("click", function () {
			body.classList.toggle("full-hex");
			set("full-hex", body.classList.contains("full-hex") ? "1" : "0");
			syncHex();
		});
	}

	// ---- quick jump ----------------------------------------------------
	var jump = document.getElementById("jump");
	document.addEventListener("keydown", function (e) {
		var tag = (e.target.tagName || "").toLowerCase();
		var typing = tag === "input" || tag === "textarea" || tag === "select";
		if (e.key === "/" && !typing && jump) { e.preventDefault(); jump.focus(); jump.select(); }
		if (e.key === "[" && !typing && navToggle) { e.preventDefault(); navToggle.click(); }
		if (e.key === "Escape") {
			var drawer = document.getElementById("json-drawer");
			if (drawer && !drawer.hidden) drawer.hidden = true;
			if (typing) e.target.blur();
		}
	});

	// ---- JSON drawer ---------------------------------------------------
	var jsonToggle = document.getElementById("json-toggle");
	var drawerEl = document.getElementById("json-drawer");
	var jdBody = document.getElementById("jd-body");
	var jdPath = document.getElementById("jd-path");
	var jdFilter = document.getElementById("jd-filter");
	var jdClose = document.getElementById("jd-close");
	var rawLines = null;

	function esc(s) {
		return s.replace(/&/g, "&amp;").replace(/</g, "&lt;").replace(/>/g, "&gt;");
	}
	function colorize(line) {
		// key: "..." / strings / numbers / booleans+null, per line.
		var out = esc(line);
		out = out.replace(/^(\s*)("(?:[^"\\]|\\.)*")(\s*:)/, '$1<span class="jk">$2</span>$3');
		out = out.replace(/(:\s*)("(?:[^"\\]|\\.)*")/, '$1<span class="js">$2</span>');
		out = out.replace(/(:\s*)(-?\d+(?:\.\d+)?(?:e[+-]?\d+)?)/i, '$1<span class="jn">$2</span>');
		out = out.replace(/(:\s*)(true|false|null)\b/, '$1<span class="jb">$2</span>');
		return out;
	}
	function renderJSON() {
		if (!rawLines) return;
		var q = (jdFilter.value || "").toLowerCase();
		var html = [];
		for (var i = 0; i < rawLines.length; i++) {
			var l = rawLines[i];
			if (q && l.toLowerCase().indexOf(q) < 0) continue;
			html.push(q ? '<span class="hit">' + colorize(l) + "</span>" : colorize(l));
		}
		jdBody.innerHTML = html.join("\n") || '<span class="dim">no matching lines</span>';
	}
	function openDrawer(api) {
		drawerEl.hidden = false;
		jdPath.textContent = api;
		jdBody.textContent = "loading " + api + " …";
		fetch(api, { headers: { Accept: "application/json" } })
			.then(function (r) { return r.text(); })
			.then(function (txt) {
				var pretty = txt;
				try { pretty = JSON.stringify(JSON.parse(txt), null, 2); } catch (e) { /* show raw */ }
				rawLines = pretty.split("\n");
				renderJSON();
			})
			.catch(function (err) { jdBody.textContent = "fetch failed: " + err; });
	}
	if (jsonToggle && drawerEl) {
		jsonToggle.addEventListener("click", function () {
			if (!drawerEl.hidden) { drawerEl.hidden = true; return; }
			openDrawer(jsonToggle.getAttribute("data-api"));
		});
		jdClose.addEventListener("click", function () { drawerEl.hidden = true; });
		jdFilter.addEventListener("input", renderJSON);
	}
})();
