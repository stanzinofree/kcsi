/**
 * Support CTA Component
 * Injects consistent support call-to-action block across documentation pages
 */

(function() {
    'use strict';
    
    // Support CTA HTML template
    const SUPPORT_CTA_HTML = `
        <div class="support-cta-block">
            <h3>☕ Support KCSI</h3>
            <p>KCSI is free and open source. If it saves you time, consider supporting:</p>
            <div class="support-buttons">
                <a href="https://buymeacoffee.com/smilzao" class="btn btn-support" target="_blank" rel="noopener">
                    ☕ Buy Me a Coffee
                </a>
                <a href="https://github.com/sponsors/stanzinofree" class="btn btn-secondary" target="_blank" rel="noopener">
                    💝 GitHub Sponsors
                </a>
                <a href="support.html" class="btn btn-primary">
                    👥 For Teams
                </a>
            </div>
            <p class="support-note">No telemetry • Open source • Audit the code yourself</p>
        </div>
    `;
    
    /**
     * Inject support CTA before footer
     */
    function injectSupportCTA() {
        const footer = document.querySelector('footer');
        if (!footer) {
            console.warn('KCSI Support CTA: Footer element not found');
            return;
        }
        
        // Create container and inject HTML
        const container = document.createElement('div');
        container.innerHTML = SUPPORT_CTA_HTML.trim();
        
        // Insert before footer
        footer.parentNode.insertBefore(container.firstChild, footer);
    }
    
    // Inject when DOM is ready
    if (document.readyState === 'loading') {
        document.addEventListener('DOMContentLoaded', injectSupportCTA);
    } else {
        injectSupportCTA();
    }
})();
