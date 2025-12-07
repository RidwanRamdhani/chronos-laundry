import { API_BASE, getToken, requireAuth, logout } from "/src/main.js";

requireAuth();  // pastikan sudah login

const adminName = document.getElementById("layoutAdminName");
const logoutBtn = document.getElementById("layoutLogoutBtn");

if (adminName) {
    const fullname = localStorage.getItem("admin_fullname") || "";
    adminName.innerHTML = `<i class="fas fa-user-circle me-2"></i>${fullname}`;
}

if (logoutBtn) {
    logoutBtn.addEventListener("click", () => logout());
}

// === Animate Number Counter ===
function animateValue(element, start, end, duration) {
    const range = end - start;
    const increment = range / (duration / 16); // 60fps
    let current = start;
    
    const timer = setInterval(() => {
        current += increment;
        if ((increment > 0 && current >= end) || (increment < 0 && current <= end)) {
            current = end;
            clearInterval(timer);
        }
        element.textContent = Math.floor(current);
    }, 16);
}

// === Show Loading State ===
function showLoading() {
    const statValues = document.querySelectorAll('.stat-value');
    statValues.forEach(el => {
        el.innerHTML = '<i class="fas fa-spinner fa-spin"></i>';
    });
}

// === Fetch Dashboard Data ===
async function loadDashboard() {
    showLoading();
    
    try {
        const response = await fetch(`${API_BASE}/transactions/dashboard`, {
            method: "GET",
            headers: {
                "Authorization": "Bearer " + getToken()
            }
        });

        const result = await response.json();

        if (!response.ok) {
            throw new Error("Gagal mengambil data dashboard");
        }

        const data = result.data;

        // Animate each counter with delay
        setTimeout(() => {
            animateValue(document.getElementById("total"), 0, data.total || 0, 1000);
        }, 100);
        
        setTimeout(() => {
            animateValue(document.getElementById("antrian"), 0, data.antrian || 0, 1000);
        }, 200);
        
        setTimeout(() => {
            animateValue(document.getElementById("mencuci"), 0, data.mencuci || 0, 1000);
        }, 300);
        
        setTimeout(() => {
            animateValue(document.getElementById("menyetrika"), 0, data.menyetrika || 0, 1000);
        }, 400);
        
        setTimeout(() => {
            animateValue(document.getElementById("siap"), 0, data.siap_diambil || 0, 1000);
        }, 500);
        
        setTimeout(() => {
            animateValue(document.getElementById("selesai"), 0, data.selesai || 0, 1000);
        }, 600);

    } catch (err) {
        console.error(err);
        
        // Show error state
        const statValues = document.querySelectorAll('.stat-value');
        statValues.forEach(el => {
            el.textContent = '0';
        });
        
        // Show error notification
        showNotification("Gagal memuat data dashboard. Silakan refresh halaman.", "error");
    }
}

// === Show Notification ===
function showNotification(message, type = "info") {
    const notification = document.createElement('div');
    notification.style.cssText = `
        position: fixed;
        top: 20px;
        right: 20px;
        padding: 15px 25px;
        background: ${type === 'error' ? 'linear-gradient(135deg, #f56565, #e53e3e)' : 'linear-gradient(135deg, #667eea, #764ba2)'};
        color: white;
        border-radius: 10px;
        box-shadow: 0 10px 30px rgba(0,0,0,0.2);
        z-index: 9999;
        animation: slideInRight 0.3s ease-out;
        font-weight: 500;
    `;
    notification.textContent = message;
    
    document.body.appendChild(notification);
    
    setTimeout(() => {
        notification.style.animation = 'slideOutRight 0.3s ease-out';
        setTimeout(() => notification.remove(), 300);
    }, 3000);
}

// Add animation keyframes
const style = document.createElement('style');
style.textContent = `
    @keyframes slideInRight {
        from {
            transform: translateX(100%);
            opacity: 0;
        }
        to {
            transform: translateX(0);
            opacity: 1;
        }
    }
    
    @keyframes slideOutRight {
        from {
            transform: translateX(0);
            opacity: 1;
        }
        to {
            transform: translateX(100%);
            opacity: 0;
        }
    }
`;
document.head.appendChild(style);

// === Auto Refresh Dashboard ===
let refreshInterval;

function startAutoRefresh() {
    // Refresh every 30 seconds
    refreshInterval = setInterval(() => {
        loadDashboard();
    }, 30000);
}

function stopAutoRefresh() {
    if (refreshInterval) {
        clearInterval(refreshInterval);
    }
}

// Stop auto-refresh when page is hidden
document.addEventListener('visibilitychange', () => {
    if (document.hidden) {
        stopAutoRefresh();
    } else {
        startAutoRefresh();
    }
});

// Initial load
loadDashboard();
startAutoRefresh();

// Cleanup on page unload
window.addEventListener('beforeunload', () => {
    stopAutoRefresh();
});
