import { requireAuth, logout } from "/src/main.js";

// Pastikan user sudah login
requireAuth();

// Ambil container layout
const layoutContainer = document.getElementById("layout");

// Inject layout from file
async function loadLayout() {
    const layoutHtml = await fetch("/layouts/layout.html").then(r => r.text());
    layoutContainer.innerHTML = layoutHtml;

    setupLayoutFunctions();
    highlightActiveMenu();

    // ==== PEMINDAHAN KONTEN HARUS DI SINI ====
    const pageContent = document.getElementById("pageContent");
    const layoutMain = document.getElementById("layoutMain");

    if (pageContent && layoutMain) {
        // Show the content and move it to layoutMain
        pageContent.style.display = 'block';
        layoutMain.appendChild(pageContent);
    }
}


loadLayout();

// ================================
// Setup events (logout, sidebar)
// ================================
function setupLayoutFunctions() {
    document.getElementById("layoutAdminName").textContent =
        localStorage.getItem("admin_fullname") || "";

    document.getElementById("layoutLogoutBtn").addEventListener("click", () => logout());

    document.getElementById("toggleSidebar").addEventListener("click", () => {
        const sidebar = document.getElementById("sidebar");
        sidebar.classList.toggle("collapsed");
    });
}

// ================================
// Highlight menu aktif
// ================================
function highlightActiveMenu() {
    const currentPath = window.location.pathname;

    document.querySelectorAll("#sidebarMenu .nav-link").forEach((link) => {
        if (link.getAttribute("href") === currentPath) {
            link.classList.add("active", "bg-primary", "text-white");
        }
    });
}
