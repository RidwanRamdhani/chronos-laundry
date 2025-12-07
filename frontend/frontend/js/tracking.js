const API_BASE = "http://localhost:8080/api";

document.getElementById("btnTrack").addEventListener("click", fetchTracking);

async function fetchTracking() {
    const code = document.getElementById("trackCode").value.trim();
    const errorBox = document.getElementById("trackError");
    const btnTrack = document.getElementById("btnTrack");

    errorBox.classList.add("d-none");
    errorBox.innerHTML = "";

    if (!code) {
        showError("Kode transaksi tidak boleh kosong.");
        return;
    }

    // Show loading state
    btnTrack.disabled = true;
    btnTrack.innerHTML = '<span class="loading-spinner"></span> Mencari...';

    try {
        const res = await fetch(`${API_BASE}/track/${code}`);

        if (!res.ok) {
            showError("Kode transaksi tidak ditemukan. Periksa kembali kode Anda.");
            return;
        }

        const json = await res.json();
        renderTracking(json.data);

    } catch (err) {
        console.error(err);
        showError("Gagal menghubungi server. Pastikan koneksi internet Anda stabil.");
    } finally {
        // Reset button
        btnTrack.disabled = false;
        btnTrack.innerHTML = '<i class="fas fa-location-arrow"></i> Lacak Sekarang';
    }
}

function showError(message) {
    const errorBox = document.getElementById("trackError");
    errorBox.innerHTML = `
        <div class="error-message">
            <i class="fas fa-exclamation-circle"></i>
            ${message}
        </div>
    `;
    errorBox.classList.remove("d-none");
    
    // Hide result if visible
    document.getElementById("trackingResult").classList.add("d-none");
}

function renderTracking(data) {
    document.getElementById("trackingResult").classList.remove("d-none");

    // Basic info
    document.getElementById("invoiceDisplay").textContent = data.transaction_code;
    document.getElementById("customerName").textContent = data.customer_name;
    document.getElementById("itemsCount").textContent = data.items_count + " item";
    document.getElementById("totalPrice").textContent = "Rp " + data.total_price.toLocaleString("id-ID");
    
    // Payment status with icon
    const isPaidEl = document.getElementById("isPaid");
    if (data.is_paid) {
        isPaidEl.innerHTML = '<i class="fas fa-check-circle" style="color: #28a745;"></i> Lunas';
    } else {
        isPaidEl.innerHTML = '<i class="fas fa-times-circle" style="color: #dc3545;"></i> Belum Lunas';
    }
    
    // Format pickup date to be more readable
    if (data.pickup_date) {
        const pickupDate = new Date(data.pickup_date);
        const formattedDate = pickupDate.toLocaleDateString("id-ID", {
            weekday: 'long',
            day: 'numeric',
            month: 'long',
            year: 'numeric'
        });
        document.getElementById("eta").textContent = formattedDate;
    } else {
        document.getElementById("eta").textContent = "Belum ditentukan";
    }

    // Status badge
    renderStatusBadge(data.status);

    // Progress steps
    highlightSteps(data.status);

    // Timeline
    renderTimeline(data.status_history);

    // Scroll to result
    setTimeout(() => {
        document.getElementById("trackingResult").scrollIntoView({ 
            behavior: 'smooth', 
            block: 'start' 
        });
    }, 100);
}

function renderStatusBadge(status) {
    const statusBadge = document.getElementById("statusBadge");
    
    const statusConfig = {
        'Queued': { icon: 'clock', text: 'Queued' },
        'Washing': { icon: 'water', text: 'Washing' },
        'Ironing': { icon: 'fire', text: 'Ironing' },
        'Ready to Pick Up': { icon: 'box', text: 'Ready to Pick Up' },
        'Complete': { icon: 'check-circle', text: 'Completed' }
    };

    const config = statusConfig[status] || { icon: 'info-circle', text: status };
    
    statusBadge.innerHTML = `
        <div class="status-badge ${status.replace(/\s+/g, '-')}">
            <i class="fas fa-${config.icon}"></i>
            ${config.text}
        </div>
    `;
}

function highlightSteps(currentStatus) {
    const steps = ["Queued", "Washing", "Ironing", "Ready to Pick Up", "Complete"];
    const currentIndex = steps.indexOf(currentStatus);

    // If status not found, default to first step
    if (currentIndex === -1) {
        console.warn(`Status "${currentStatus}" not found in steps array`);
        return;
    }

    steps.forEach((step, index) => {
        const stepEl = document.querySelector(`.step-item[data-step="${step}"]`);
        
        if (!stepEl) {
            console.warn(`Step element not found for: ${step}`);
            return;
        }
        
        // Remove all classes first
        stepEl.classList.remove("active", "completed");
        
        if (index < currentIndex) {
            // Completed steps
            stepEl.classList.add("completed");
        } else if (index === currentIndex) {
            // Current step
            stepEl.classList.add("active");
        }
    });

    // Animate progress line
    const progressLine = document.getElementById("progressLine");
    
    if (!progressLine) {
        console.warn("Progress line element not found");
        return;
    }
    
    // Calculate progress percentage
    // For 5 steps (0-4 indices), we want:
    // Index 0 (antrian) = 0%
    // Index 1 (mencuci) = 25%
    // Index 2 (menyetrika) = 50%
    // Index 3 (siap_diambil) = 75%
    // Index 4 (selesai) = 100%
    const progressPercentage = (currentIndex / (steps.length - 1)) * 100;
    
    // Reset to 0 first
    progressLine.style.width = "0%";
    
    // Force a reflow to ensure the reset is applied
    progressLine.offsetHeight;
    
    // Animate to target width after a small delay
    setTimeout(() => {
        progressLine.style.width = progressPercentage + "%";
    }, 100);
}

function renderTimeline(history) {
    const container = document.getElementById("statusHistory");
    container.innerHTML = "";

    if (!history || history.length === 0) {
        container.innerHTML = `
            <div style="text-align: center; padding: 40px; color: #999;">
                <i class="fas fa-inbox" style="font-size: 3rem; margin-bottom: 15px; opacity: 0.5;"></i>
                <p style="margin: 0;">Belum ada riwayat perubahan status.</p>
            </div>
        `;
        return;
    }

    // Sort history by date (newest first)
    const sortedHistory = [...history].sort((a, b) => 
        new Date(b.created_at) - new Date(a.created_at)
    );

    sortedHistory.forEach((item, index) => {
        const time = new Date(item.created_at).toLocaleString("id-ID", {
            day: '2-digit',
            month: 'long',
            year: 'numeric',
            hour: '2-digit',
            minute: '2-digit'
        });

        const timelineItem = document.createElement('div');
        timelineItem.className = 'timeline-item';
        timelineItem.style.animationDelay = `${index * 0.1}s`;

        timelineItem.innerHTML = `
            <div class="timeline-content">
                <div class="timeline-status">
                    <i class="fas fa-arrow-right" style="color: #667eea; font-size: 0.9rem;"></i>
                    ${formatStatus(item.previous_status)} → ${formatStatus(item.new_status)}
                </div>
                <div class="timeline-meta">
                    <span>
                        <i class="fas fa-calendar-alt"></i>
                        ${time}
                    </span>
                    <span>
                        <i class="fas fa-user"></i>
                        ${item.changed_by}
                    </span>
                </div>
                ${item.reason ? `
                    <div class="timeline-reason">
                        <i class="fas fa-comment-dots"></i>
                        ${item.reason}
                    </div>
                ` : ''}
            </div>
            ${index < sortedHistory.length - 1 ? '<div class="timeline-line"></div>' : ''}
        `;

        container.appendChild(timelineItem);
    });
}

function formatStatus(status) {
    // Status already in English from API, just return it
    return status;
}
