/**
 * KCSI Unified UI Components
 * Injects header, footer, and support callout across all documentation pages
 * Single source of truth for brand consistency and monetization
 */

(function() {
    'use strict';
    
    // ============================================================================
    // CONFIGURATION
    // ============================================================================
    
    const KCSI_CONFIG = {
        brand: {
            name: 'KCSI',
            tagline: 'kubectl for humans — Cascading TAB + guardrails for day-2 ops',
            description: 'Cascading TAB autocomplete + guardrails for Kubernetes day-2 operations'
        },
        
        nav: [
            { label: 'Home', href: 'index.html', icon: '🏠' },
            { label: 'Cheatsheet', href: 'cheatsheet.html', icon: '📖' },
            { label: 'Roadmap', href: 'roadmap.html', icon: '🗺️' },
            { label: 'For Teams', href: 'teams.html', icon: '👥' },
            { label: 'Support', href: 'support.html', icon: '☕', highlight: true }
        ],
        
        links: {
            github: 'https://github.com/stanzinofree/kcsi',
            issues: 'https://github.com/stanzinofree/kcsi/issues',
            license: 'https://github.com/stanzinofree/kcsi/blob/main/LICENSE',
            coffee: 'https://buymeacoffee.com/smilzao',
            sponsors: 'https://github.com/sponsors/stanzinofree',
            cv: 'https://cv.middei.info/',
            linkedin: 'https://www.linkedin.com/in/stanzinofree/',
            githubProfile: 'https://github.com/stanzinofree'
        },
        
        author: {
            name: 'Alessandro Middei'
        },
        
        trustSignal: 'No telemetry by default • Open source • Audit the code yourself'
    };
    
    // ============================================================================
    // UTILITIES
    // ============================================================================
    
    /**
     * Get current page filename from URL
     */
    function getCurrentPage() {
        const path = window.location.pathname;
        const filename = path.substring(path.lastIndexOf('/') + 1);
        return filename || 'index.html';
    }
    
    /**
     * Check if nav item should be marked as active
     */
    function isActivePage(href) {
        const current = getCurrentPage();
        return current === href;
    }
    
    // ============================================================================
    // HEADER COMPONENT
    // ============================================================================
    
    function renderHeader() {
        const navItems = KCSI_CONFIG.nav.map(item => {
            const activeClass = isActivePage(item.href) ? 'nav-active' : '';
            const highlightClass = item.highlight ? 'nav-highlight' : '';
            return `
                <a href="${item.href}" class="nav-item ${activeClass} ${highlightClass}">
                    ${item.icon} ${item.label}
                </a>
            `;
        }).join('');
        
        return `
            <header class="kcsi-header">
                <div class="kcsi-header-container">
                    <div class="kcsi-brand">
                        <a href="index.html" class="kcsi-logo">
                            <strong>${KCSI_CONFIG.brand.name}</strong>
                        </a>
                        <p class="kcsi-tagline">${KCSI_CONFIG.brand.tagline}</p>
                    </div>
                    <nav class="kcsi-nav">
                        ${navItems}
                    </nav>
                </div>
            </header>
        `;
    }
    
    // ============================================================================
    // FOOTER COMPONENT
    // ============================================================================
    
    function renderFooter() {
        return `
            <footer class="kcsi-footer">
                <div class="kcsi-footer-content">
                    <div class="kcsi-footer-brand">
                        <h3>${KCSI_CONFIG.brand.name} - kubectl for humans</h3>
                        <p>${KCSI_CONFIG.brand.description}</p>
                    </div>
                    
                    <div class="kcsi-footer-links">
                        <a href="${KCSI_CONFIG.links.github}" target="_blank" rel="noopener">
                            <svg width="20" height="20" viewBox="0 0 16 16" fill="currentColor">
                                <path d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.013 8.013 0 0016 8c0-4.42-3.58-8-8-8z"></path>
                            </svg>
                            GitHub
                        </a>
                        <a href="${KCSI_CONFIG.links.issues}" target="_blank" rel="noopener">Report Issue</a>
                        <a href="${KCSI_CONFIG.links.license}" target="_blank" rel="noopener">MIT License</a>
                        <a href="${KCSI_CONFIG.links.coffee}" target="_blank" rel="noopener">☕ Support</a>
                    </div>
                    
                    <div class="kcsi-footer-author">
                        <p>Created by <strong>${KCSI_CONFIG.author.name}</strong></p>
                        <div class="kcsi-author-links">
                            <a href="${KCSI_CONFIG.links.cv}" target="_blank" rel="noopener">📄 Resume/CV</a>
                            <a href="${KCSI_CONFIG.links.linkedin}" target="_blank" rel="noopener">💼 LinkedIn</a>
                            <a href="${KCSI_CONFIG.links.githubProfile}" target="_blank" rel="noopener">🐙 GitHub Profile</a>
                        </div>
                    </div>
                </div>
            </footer>
        `;
    }
    
    // ============================================================================
    // SUPPORT CALLOUT COMPONENT
    // ============================================================================
    
    function renderSupportCallout() {
        return `
            <div class="support-cta-block">
                <h3>☕ Support ${KCSI_CONFIG.brand.name}</h3>
                <p>${KCSI_CONFIG.brand.name} is free and open source. If it saves you time, consider supporting:</p>
                <div class="support-buttons">
                    <a href="${KCSI_CONFIG.links.coffee}" class="btn btn-support" target="_blank" rel="noopener">
                        ☕ Buy Me a Coffee
                    </a>
                    <a href="${KCSI_CONFIG.links.sponsors}" class="btn btn-secondary" target="_blank" rel="noopener">
                        💝 GitHub Sponsors
                    </a>
                    <a href="teams.html" class="btn btn-primary">
                        👥 For Teams
                    </a>
                </div>
                <p class="support-note">${KCSI_CONFIG.trustSignal}</p>
            </div>
        `;
    }
    
    // ============================================================================
    // INJECTION FUNCTIONS
    // ============================================================================
    
    /**
     * Inject header into #kcsi-header placeholder
     */
    function injectHeader() {
        const placeholder = document.getElementById('kcsi-header');
        if (placeholder) {
            placeholder.innerHTML = renderHeader();
        }
    }
    
    /**
     * Inject footer into #kcsi-footer placeholder
     */
    function injectFooter() {
        const placeholder = document.getElementById('kcsi-footer');
        if (placeholder) {
            placeholder.innerHTML = renderFooter();
        }
    }
    
    /**
     * Inject support callout into all .kcsi-support-slot placeholders
     */
    function injectSupportCallout() {
        const placeholders = document.querySelectorAll('.kcsi-support-slot, #kcsi-support');
        placeholders.forEach(placeholder => {
            placeholder.innerHTML = renderSupportCallout();
        });
    }
    
    /**
     * Main initialization function
     */
    function initKCSIUI() {
        injectHeader();
        injectFooter();
        injectSupportCallout();
    }
    
    // ============================================================================
    // AUTO-INIT
    // ============================================================================
    
    // Run when DOM is ready
    if (document.readyState === 'loading') {
        document.addEventListener('DOMContentLoaded', initKCSIUI);
    } else {
        initKCSIUI();
    }
    
    // Expose config for debugging (optional)
    window.KCSI_UI = {
        config: KCSI_CONFIG,
        reinit: initKCSIUI
    };
})();
