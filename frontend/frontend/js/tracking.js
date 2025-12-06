const API_BASE = "http://localhost:8080/api";

document.getElementById("btnTrack").addEventListener("click", fetchTracking);

async function fetchTracking() {
    const code = document.getElementById("trackCode").value.trim();
    const errorBox = document.getElementById("trackError");

    errorBox.classList.add("d-none");
    errorBox.textContent = "";

    if (!code) {
        errorBox.textContent = "Kode transaksi tidak boleh kosong.";
        errorBox.classList.remove("d-none");
        return;
    }

    try {
        const res = await fetch(`${API_BASE}/track/${code}`);

        if (!res.ok) {
            errorBox.textContent = "Kode transaksi tidak ditemukan.";
            errorBox.classList.remove("d-none");
            return;
        }

        const json = await res.json();
        renderTracking(json.data);

    } catch (err) {
        console.error(err);
        errorBox.textContent = "Gagal menghubungi server.";
        errorBox.classList.remove("d-none");
    }
}

function renderTracking(data) {
    document.getElementById("trackingResult").classList.remove("d-none");

    document.getElementById("invoiceDisplay").textContent = data.transaction_code;
    document.getElementById("customerName").textContent = data.customer_name;
    document.getElementById("currentStatus").textContent = data.status;
    document.getElementById("itemsCount").textContent = data.items_count + " item";
    document.getElementById("totalPrice").textContent = "Rp " + data.total_price.toLocaleString("id-ID");
    document.getElementById("isPaid").textContent = data.is_paid ? "Ya" : "Belum";
    document.getElementById("eta").textContent = data.pickup_date || "-";

    highlightSteps(data.status);
    renderTimeline(data.status_history);
}

function highlightSteps(currentStatus) {
    const steps = ["antrian", "mencuci", "menyetrika", "siap_diambil", "selesai"];

    steps.forEach(step => {
        const el = document.getElementById(`step-${step}`);
        el.classList.remove("active");

        if (steps.indexOf(step) <= steps.indexOf(currentStatus)) {
            el.classList.add("active");
        }
    });
}

function renderTimeline(history) {
    const container = document.getElementById("statusHistory");
    container.innerHTML = "";

    if (!history || history.length === 0) {
        container.innerHTML = `<p class="text-muted">Belum ada riwayat status.</p>`;
        return;
    }

    history.forEach((item, index) => {
        const time = new Date(item.created_at).toLocaleString("id-ID");

        container.innerHTML += `
            <div class="timeline-item">
                <div class="fw-bold">${item.previous_status} → ${item.new_status}</div>
                <div class="text-muted" style="font-size: 0.9rem;">
                    ${time} • oleh ${item.changed_by}
                </div>
                <div>${item.reason || ""}</div>

                ${index < history.length - 1 ? `<div class="timeline-line"></div>` : ""}
            </div>
        `;
    });
}
