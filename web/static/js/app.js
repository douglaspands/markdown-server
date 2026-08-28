/**
 * Markdown Viewer - Aplicação Frontend (Interatividade, QR Code, Mermaid, Temas)
 */
document.addEventListener("DOMContentLoaded", () => {
  initTheme();
  initSidebar();
  initSidebarResizer();
  initQRCodeModal();
  initMermaid();
  initCodeCopyButtons();
  initMathEquations();
  initActiveTreeLink();
});

/* =========================================================================
   1. Gerenciamento de Tema (Dark / Light Mode)
   ========================================================================= */
function initTheme() {
  const themeToggleBtn = document.getElementById("theme-toggle");
  const themeIcon = document.getElementById("theme-icon");
  const html = document.documentElement;

  const savedTheme = localStorage.getItem("md_theme") ||
    (window.matchMedia("(prefers-color-scheme: dark)").matches ? "dark" : "light");

  applyTheme(savedTheme);

  if (themeToggleBtn) {
    themeToggleBtn.addEventListener("click", () => {
      const currentTheme = html.getAttribute("data-theme") || "light";
      const newTheme = currentTheme === "dark" ? "light" : "dark";
      applyTheme(newTheme);
      localStorage.setItem("md_theme", newTheme);
      renderMermaidDiagrams(newTheme);
    });
  }

  function applyTheme(theme) {
    html.setAttribute("data-theme", theme);
    if (themeIcon) {
      themeIcon.textContent = theme === "dark" ? "🌙" : "☀️";
    }
  }
}

/* =========================================================================
   2. Menu Lateral Retrátil (Sidebar Toggle) & Suporte a Mobile Drawer
   ========================================================================= */
function initSidebar() {
  const sidebar = document.getElementById("app-sidebar");
  const toggleBtn = document.getElementById("sidebar-toggle-btn");
  const backdrop = document.getElementById("sidebar-backdrop");

  if (!sidebar || !toggleBtn) return;

  // Restaura estado no Desktop
  const isDesktop = window.innerWidth > 768;
  const isCollapsed = localStorage.getItem("md_sidebar_collapsed") === "true";

  if (isDesktop && isCollapsed) {
    sidebar.classList.add("collapsed");
  }

  // Toggle do menu
  toggleBtn.addEventListener("click", () => {
    if (window.innerWidth <= 768) {
      // Mobile / Tablet Drawer
      const isOpen = sidebar.classList.contains("mobile-open");
      if (isOpen) {
        closeMobileSidebar();
      } else {
        openMobileSidebar();
      }
    } else {
      // Desktop Collapsible
      sidebar.classList.toggle("collapsed");
      localStorage.setItem("md_sidebar_collapsed", sidebar.classList.contains("collapsed"));
    }
  });

  // Atalho de Teclado (Ctrl+B ou Cmd+B)
  document.addEventListener("keydown", (e) => {
    if ((e.ctrlKey || e.metaKey) && e.key.toLowerCase() === "b") {
      e.preventDefault();
      toggleBtn.click();
    }
  });

  // Fecha gaveta mobile ao clicar no backdrop ou em um link
  if (backdrop) {
    backdrop.addEventListener("click", closeMobileSidebar);
  }

  document.querySelectorAll(".file-link").forEach((link) => {
    link.addEventListener("click", () => {
      if (window.innerWidth <= 768) {
        closeMobileSidebar();
      }
    });
  });

  function openMobileSidebar() {
    sidebar.classList.add("mobile-open");
    if (backdrop) backdrop.classList.add("active");
  }

  function closeMobileSidebar() {
    sidebar.classList.remove("mobile-open");
    if (backdrop) backdrop.classList.remove("active");
  }

  // Toggle de Diretórios na Árvore
  document.querySelectorAll(".dir-title").forEach((title) => {
    title.addEventListener("click", (e) => {
      e.stopPropagation();
      const parent = title.parentElement;
      const children = parent.querySelector(".dir-children");
      const arrow = title.querySelector(".dir-arrow");

      if (children) {
        const isHidden = children.style.display === "none";
        children.style.display = isHidden ? "block" : "none";
        if (arrow) {
          arrow.textContent = isHidden ? "▾" : "▸";
        }
      }
    });
  });
}

/* =========================================================================
   3. Divisor Maleável de Redimensionamento da Barra Lateral (Splitter)
   ========================================================================= */
function initSidebarResizer() {
  const resizer = document.getElementById("sidebar-resizer");
  const sidebar = document.getElementById("app-sidebar");
  if (!resizer || !sidebar) return;

  let isDragging = false;
  const MIN_WIDTH = 180;

  // Restaura largura previamente customizada pelo usuário
  const savedWidth = localStorage.getItem("md_sidebar_width");
  if (savedWidth && window.innerWidth > 768) {
    const parsed = parseInt(savedWidth, 10);
    if (!isNaN(parsed) && parsed >= MIN_WIDTH) {
      sidebar.style.width = `${parsed}px`;
      sidebar.style.minWidth = `${parsed}px`;
    }
  }

  const onMouseDown = (e) => {
    if (window.innerWidth <= 768 || sidebar.classList.contains("collapsed")) return;
    isDragging = true;
    document.body.classList.add("resizing");
    document.addEventListener("mousemove", onMouseMove);
    document.addEventListener("mouseup", onMouseUp);
    e.preventDefault();
  };

  const onMouseMove = (e) => {
    if (!isDragging) return;
    const maxWidth = Math.min(600, Math.floor(window.innerWidth * 0.55));
    let newWidth = e.clientX;
    if (newWidth < MIN_WIDTH) newWidth = MIN_WIDTH;
    if (newWidth > maxWidth) newWidth = maxWidth;

    sidebar.style.width = `${newWidth}px`;
    sidebar.style.minWidth = `${newWidth}px`;
  };

  const onMouseUp = () => {
    if (!isDragging) return;
    isDragging = false;
    document.body.classList.remove("resizing");
    document.removeEventListener("mousemove", onMouseMove);
    document.removeEventListener("mouseup", onMouseUp);

    const currentWidth = parseInt(sidebar.style.width, 10);
    if (!isNaN(currentWidth) && currentWidth >= MIN_WIDTH) {
      localStorage.setItem("md_sidebar_width", currentWidth);
    }
  };

  resizer.addEventListener("mousedown", onMouseDown);
}

/* =========================================================================
   4. Modal e Geração de QR Code com Endereço de Rede Local (LAN)
   ========================================================================= */
function initQRCodeModal() {
  const qrBtn = document.getElementById("qrcode-btn");
  const modal = document.getElementById("qrcode-modal");
  const closeBtn = document.getElementById("close-qrcode-modal");
  const qrContainer = document.getElementById("qrcode-canvas");
  const urlInput = document.getElementById("share-url-input");
  const copyBtn = document.getElementById("copy-url-btn");

  if (!qrBtn || !modal) return;

  qrBtn.addEventListener("click", () => {
    // Determina a URL acessível na rede local
    const lanBaseURL = modal.getAttribute("data-lan-url");
    let shareURL = "";

    if (lanBaseURL && !lanBaseURL.includes("127.0.0.1") && !lanBaseURL.includes("localhost")) {
      try {
        const lanUrlObj = new URL(lanBaseURL);
        shareURL = `${lanUrlObj.protocol}//${lanUrlObj.host}${window.location.pathname}${window.location.search}${window.location.hash}`;
      } catch (err) {
        shareURL = `${lanBaseURL.replace(/\/$/, "")}${window.location.pathname}${window.location.search}${window.location.hash}`;
      }
    } else {
      shareURL = window.location.href;
    }

    if (urlInput) urlInput.value = shareURL;

    if (qrContainer && window.QRCode) {
      qrContainer.innerHTML = "";
      new window.QRCode(qrContainer, {
        text: shareURL,
        width: 190,
        height: 190
      });
    }

    modal.classList.add("active");
  });

  if (closeBtn) {
    closeBtn.addEventListener("click", () => {
      modal.classList.remove("active");
    });
  }

  modal.addEventListener("click", (e) => {
    if (e.target === modal) {
      modal.classList.remove("active");
    }
  });

  if (copyBtn && urlInput) {
    copyBtn.addEventListener("click", () => {
      urlInput.select();
      navigator.clipboard.writeText(urlInput.value).then(() => {
        const originalText = copyBtn.textContent;
        copyBtn.textContent = "Copiado!";
        setTimeout(() => {
          copyBtn.textContent = originalText;
        }, 1500);
      });
    });
  }
}

/* =========================================================================
   5. Renderização Assíncrona e Robusta de Diagramas Mermaid.js
   ========================================================================= */
function initMermaid() {
  const mermaidBlocks = document.querySelectorAll(".mermaid");
  if (mermaidBlocks.length === 0) return;

  // Salva o código-fonte original em data-original-code para viabilizar re-renderizações
  mermaidBlocks.forEach((el, idx) => {
    if (!el.getAttribute("data-original-code")) {
      el.setAttribute("data-original-code", el.textContent.trim());
      el.setAttribute("id", `mermaid-block-${idx}`);
    }
  });

  const currentTheme = document.documentElement.getAttribute("data-theme") || "light";
  renderMermaidDiagrams(currentTheme);
}

async function renderMermaidDiagrams(theme) {
  const mermaidBlocks = document.querySelectorAll(".mermaid");
  if (mermaidBlocks.length === 0) return;

  if (typeof window.loadMermaidLibrary === "function") {
    try {
      await window.loadMermaidLibrary();
    } catch (err) {
      console.warn("Não foi possível carregar o módulo Mermaid:", err);
      return;
    }
  }

  if (!window.mermaid) return;

  try {
    window.mermaid.initialize({
      startOnLoad: false,
      theme: theme === "dark" ? "dark" : "default",
      securityLevel: "loose",
      fontFamily: "inherit",
      logLevel: "error"
    });

    for (let i = 0; i < mermaidBlocks.length; i++) {
      const el = mermaidBlocks[i];
      const rawCode = el.getAttribute("data-original-code");
      if (!rawCode) continue;

      const containerId = `mermaid-svg-${i}-${Date.now()}`;
      try {
        const { svg } = await window.mermaid.render(containerId, rawCode);
        el.innerHTML = svg;
        el.classList.remove("mermaid-error");
        const svgEl = el.querySelector("svg");
        if (svgEl) {
          attachMermaidPanZoom(el, svgEl);
        }
      } catch (renderErr) {
        console.warn(`Erro de sintaxe no diagrama Mermaid #${i}:`, renderErr);
        el.classList.add("mermaid-error");
        el.innerHTML = `
          <div class="mermaid-error-box">
            <div class="mermaid-error-header">
              <span>⚠️ Erro de sintaxe no diagrama Mermaid</span>
            </div>
            <pre class="mermaid-error-code"><code>${escapeHTML(rawCode)}</code></pre>
          </div>
        `;
        const badEl = document.getElementById(containerId);
        if (badEl) badEl.remove();
        const badD = document.getElementById("d" + containerId);
        if (badD) badD.remove();
      }
    }
  } catch (err) {
    console.warn("Erro geral na renderização Mermaid:", err);
  }
}

/**
 * Gerenciador nativo e fluido de Pan & Zoom para diagramas Mermaid (SVG)
 */
function attachMermaidPanZoom(container, svgEl) {
  if (!container || !svgEl) return;

  // Remove toolbar anterior se já existir
  const oldToolbar = container.querySelector(".mermaid-toolbar");
  if (oldToolbar) oldToolbar.remove();

  // Cria a toolbar flutuante
  const toolbar = document.createElement("div");
  toolbar.className = "mermaid-toolbar";
  toolbar.innerHTML = `
    <button class="mermaid-tool-btn" data-action="zoom-in" title="Aumentar Zoom (+)" aria-label="Aumentar Zoom">
      <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="11" cy="11" r="8"></circle><line x1="21" y1="21" x2="16.65" y2="16.65"></line><line x1="11" y1="8" x2="11" y2="14"></line><line x1="8" y1="11" x2="14" y2="11"></line></svg>
    </button>
    <button class="mermaid-tool-btn" data-action="zoom-out" title="Diminuir Zoom (-)" aria-label="Diminuir Zoom">
      <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="11" cy="11" r="8"></circle><line x1="21" y1="21" x2="16.65" y2="16.65"></line><line x1="8" y1="11" x2="14" y2="11"></line></svg>
    </button>
    <button class="mermaid-tool-btn" data-action="reset" title="Redefinir Zoom (100%)" aria-label="Redefinir">
      <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M3 12a9 9 0 1 0 9-9 9.75 9.75 0 0 0-6.74 2.74L3 8"></path><path d="M3 3v5h5"></path></svg>
    </button>
  `;
  container.appendChild(toolbar);

  let scale = 1.0;
  let translateX = 0;
  let translateY = 0;
  let isDragging = false;
  let startX = 0;
  let startY = 0;

  function updateTransform() {
    svgEl.style.transform = `translate(${translateX}px, ${translateY}px) scale(${scale})`;
  }

  // Controles de zoom da toolbar
  toolbar.querySelector('[data-action="zoom-in"]').addEventListener("click", (e) => {
    e.stopPropagation();
    scale = Math.min(5.0, scale * 1.25);
    updateTransform();
  });

  toolbar.querySelector('[data-action="zoom-out"]').addEventListener("click", (e) => {
    e.stopPropagation();
    scale = Math.max(0.4, scale / 1.25);
    updateTransform();
  });

  toolbar.querySelector('[data-action="reset"]').addEventListener("click", (e) => {
    e.stopPropagation();
    scale = 1.0;
    translateX = 0;
    translateY = 0;
    updateTransform();
  });

  // Zoom suave com roda do mouse sobre o container
  container.addEventListener("wheel", (e) => {
    e.preventDefault();
    const delta = e.deltaY < 0 ? 1.15 : 0.87;
    scale = Math.max(0.4, Math.min(5.0, scale * delta));
    updateTransform();
  }, { passive: false });

  // Pan por clique e arraste com Pointer Events
  container.addEventListener("pointerdown", (e) => {
    if (e.target.closest(".mermaid-toolbar")) return;
    isDragging = true;
    startX = e.clientX - translateX;
    startY = e.clientY - translateY;
    container.classList.add("is-panning");
    container.setPointerCapture(e.pointerId);
  });

  container.addEventListener("pointermove", (e) => {
    if (!isDragging) return;
    translateX = e.clientX - startX;
    translateY = e.clientY - startY;
    updateTransform();
  });

  const endDrag = (e) => {
    if (isDragging) {
      isDragging = false;
      container.classList.remove("is-panning");
      try {
        container.releasePointerCapture(e.pointerId);
      } catch (_) {}
    }
  };

  container.addEventListener("pointerup", endDrag);
  container.addEventListener("pointercancel", endDrag);
}

function escapeHTML(str) {
  return str
    .replace(/&/g, "&amp;")
    .replace(/</g, "&lt;")
    .replace(/>/g, "&gt;")
    .replace(/"/g, "&quot;")
    .replace(/'/g, "&#039;");
}

/* =========================================================================
   6. Botão de Cópia em Blocos de Código (Copy Code Button)
   ========================================================================= */
function initCodeCopyButtons() {
  const codeBlocks = document.querySelectorAll(".highlight, pre:not(.mermaid-error-code)");

  codeBlocks.forEach((block) => {
    // Evita duplicar o botão se já foi injetado
    if (block.querySelector(".copy-code-btn") || block.classList.contains("mermaid")) return;

    const copyBtn = document.createElement("button");
    copyBtn.className = "copy-code-btn";
    copyBtn.type = "button";
    copyBtn.setAttribute("aria-label", "Copiar código");
    copyBtn.textContent = "Copiar";

    copyBtn.addEventListener("click", (e) => {
      e.stopPropagation();
      const codeEl = block.querySelector("code") || block;
      const textToCopy = codeEl.innerText || codeEl.textContent;

      navigator.clipboard.writeText(textToCopy).then(() => {
        copyBtn.textContent = "Copiado!";
        copyBtn.classList.add("copied");

        setTimeout(() => {
          copyBtn.textContent = "Copiar";
          copyBtn.classList.remove("copied");
        }, 1800);
      }).catch((err) => {
        console.warn("Falha ao copiar código:", err);
      });
    });

    block.appendChild(copyBtn);
  });
}

/* =========================================================================
   7. Renderização Automática de Fórmulas Matemáticas LaTeX (KaTeX)
   ========================================================================= */
function initMathEquations() {
  const content = document.querySelector(".markdown-body");
  if (!content) return;

  function renderMath() {
    if (typeof renderMathInElement === "function") {
      try {
        renderMathInElement(content, {
          delimiters: [
            { left: "$$", right: "$$", display: true },
            { left: "$", right: "$", display: false },
            { left: "\\(", right: "\\)", display: false },
            { left: "\\[", right: "\\]", display: true }
          ],
          throwOnError: false
        });
      } catch (err) {
        console.warn("Erro ao processar fórmulas KaTeX:", err);
      }
    }
  }

  if (typeof renderMathInElement === "function") {
    renderMath();
  } else {
    window.addEventListener("load", renderMath);
  }
}

/* =========================================================================
   8. Destaque do Link Ativo na Árvore de Arquivos
   ========================================================================= */
function initActiveTreeLink() {
  const currentPath = window.location.pathname;
  document.querySelectorAll(".file-link").forEach((link) => {
    const href = link.getAttribute("href");
    if (href === currentPath || (currentPath === "/" && (href === "/README.md" || href === "/readme.md"))) {
      link.classList.add("active");
    }
  });
}
