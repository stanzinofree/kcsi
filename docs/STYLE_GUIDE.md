# KCSI Documentation Style Guide

**Version:** 1.0  
**Last Updated:** 2026-01-05

This document defines the visual identity and brand system for KCSI documentation.

---

## Brand Positioning

**Tagline:** kubectl for humans  
**Mission:** Cascading TAB + guardrails for day-2 ops  
**Audience:** Sysadmins, DevOps engineers, and system engineers who use Kubernetes intermittently

**Personality:**
- **Terminal-native but premium** - Feels like a tool made by engineers, for engineers
- **Confident without arrogance** - Expert guidance, zero condescension
- **Practical over theoretical** - Show, don't lecture

---

## Typography

### Fonts

**Headings & Body:**
```css
font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
```
- Source: Google Fonts
- Weight range: 400 (Regular), 500 (Medium), 700 (Bold)

**Code & Terminal:**
```css
font-family: 'JetBrains Mono', 'Monaco', 'Courier New', monospace;
```
- Source: Google Fonts
- Weight: 400 (Regular), 500 (Medium)

### Type Scale

| Element | Size | Weight | Line Height | Usage |
|---------|------|--------|-------------|-------|
| H1 | 2.5rem (40px) | 700 | 1.2 | Page titles only |
| H2 | 1.875rem (30px) | 700 | 1.3 | Section headings |
| H3 | 1.5rem (24px) | 600 | 1.4 | Subsection headings |
| H4 | 1.25rem (20px) | 600 | 1.4 | Minor headings |
| Body | 1rem (16px) | 400 | 1.6 | Paragraph text |
| Small | 0.875rem (14px) | 400 | 1.5 | Captions, metadata |
| Code Inline | 0.9em | 400 | inherit | `inline code` |
| Code Block | 0.875rem (14px) | 400 | 1.6 | Multi-line code |

---

## Color Palette

### Primary Colors

```css
--kcsi-primary: #4a5a9a;        /* Main brand blue */
--kcsi-primary-dark: #3a4a7a;   /* Hover/active states */
--kcsi-primary-light: #667eea;  /* Accents */
```

### Surface & Background

```css
--kcsi-bg-body: #f5f7fa;        /* Page background */
--kcsi-bg-surface: #ffffff;     /* Cards, containers */
--kcsi-bg-elevated: #ffffff;    /* Modals, dropdowns */
--kcsi-bg-code: #2d2d2d;        /* Code blocks */
```

### Text Colors

```css
--kcsi-text-primary: #212529;   /* Main content */
--kcsi-text-secondary: #495057; /* Supporting text */
--kcsi-text-muted: #6c757d;     /* Captions, metadata */
--kcsi-text-inverse: #ffffff;   /* On dark backgrounds */
--kcsi-text-code: #50fa7b;      /* Inline code */
```

### Semantic Colors

```css
--kcsi-success: #28a745;        /* Success states */
--kcsi-info: #17a2b8;           /* Info callouts */
--kcsi-warning: #ffc107;        /* Warning callouts */
--kcsi-danger: #dc3545;         /* Danger/error states */
```

### Gradients

```css
--kcsi-gradient-hero: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
--kcsi-gradient-footer: linear-gradient(135deg, #a8b5d1 0%, #c3dfe0 50%, #b4d4d3 100%);
```

---

## Spacing System

**Base unit:** 8px

```css
--space-xs: 0.5rem;   /* 8px */
--space-sm: 1rem;     /* 16px */
--space-md: 1.5rem;   /* 24px */
--space-lg: 2rem;     /* 32px */
--space-xl: 3rem;     /* 48px */
--space-2xl: 4rem;    /* 64px */
--space-3xl: 6rem;    /* 96px */
```

**Content Width:**
- Maximum: 1200px (readable content)
- Wide layout: 1400px (dashboards, tables)
- Narrow: 720px (articles, documentation)

---

## UI Components

### Buttons

**Primary Button:**
```css
.btn-primary {
    padding: 12px 32px;
    background: var(--kcsi-primary);
    color: white;
    border: none;
    border-radius: 24px;
    font-weight: 600;
    font-size: 1rem;
    cursor: pointer;
    transition: all 0.3s ease;
}

.btn-primary:hover {
    background: var(--kcsi-primary-dark);
    transform: translateY(-2px);
    box-shadow: 0 8px 20px rgba(74, 90, 154, 0.3);
}
```

**Secondary Button:**
```css
.btn-secondary {
    padding: 12px 32px;
    background: transparent;
    color: var(--kcsi-primary);
    border: 2px solid var(--kcsi-primary);
    border-radius: 24px;
    font-weight: 600;
    font-size: 1rem;
    cursor: pointer;
    transition: all 0.3s ease;
}

.btn-secondary:hover {
    background: var(--kcsi-primary);
    color: white;
}
```

### Callout Boxes

**Tip:**
```css
.callout-tip {
    padding: 16px 20px;
    background: #e8f5e9;
    border-left: 4px solid #28a745;
    border-radius: 4px;
    margin: 16px 0;
}
```

**Note:**
```css
.callout-note {
    padding: 16px 20px;
    background: #e7f3ff;
    border-left: 4px solid #17a2b8;
    border-radius: 4px;
    margin: 16px 0;
}
```

**Warning:**
```css
.callout-warning {
    padding: 16px 20px;
    background: #fff8e1;
    border-left: 4px solid #ffc107;
    border-radius: 4px;
    margin: 16px 0;
}
```

**Danger:**
```css
.callout-danger {
    padding: 16px 20px;
    background: #ffebee;
    border-left: 4px solid #dc3545;
    border-radius: 4px;
    margin: 16px 0;
}
```

### Code Blocks

```css
.code-block {
    position: relative;
    background: #2d2d2d;
    padding: 20px;
    border-radius: 8px;
    overflow-x: auto;
    margin: 16px 0;
}

.code-block code {
    font-family: 'JetBrains Mono', monospace;
    font-size: 0.875rem;
    line-height: 1.6;
    color: #f8f8f2;
}

/* Inline code */
code {
    background: #2d2d2d;
    color: #50fa7b;
    padding: 2px 6px;
    border-radius: 4px;
    font-family: 'JetBrains Mono', monospace;
    font-size: 0.9em;
}
```

---

## Voice & Tone Guidelines

### Writing Principles

1. **Short & Scannable**
   - Use bullet points liberally
   - Keep paragraphs to 2-3 sentences max
   - Lead with the action or benefit

2. **Confident & Practical**
   ✅ "Install KCSI in 10 seconds"  
   ❌ "KCSI might help you install things faster"

3. **Zero Jargon (unless necessary)**
   ✅ "TAB autocomplete for namespaces"  
   ❌ "Leveraging cascading selection paradigms"

4. **Empathetic to Intermittent Users**
   ✅ "You know Kubernetes. You just don't remember the exact syntax."  
   ❌ "For advanced Kubernetes power users"

### Examples

**Good:**
> KCSI eliminates context switching. Press TAB, select your namespace, done. No flags to memorize.

**Bad:**
> KCSI is a comprehensive solution that facilitates enhanced operational workflows through intelligent autocompletion mechanisms.

---

## Accessibility

- **Contrast Ratios:** Minimum 4.5:1 for normal text, 3:1 for large text
- **Focus States:** All interactive elements must have visible focus indicator
- **Alt Text:** All images require descriptive alt attributes
- **Semantic HTML:** Use proper heading hierarchy, landmarks, ARIA when needed

---

## File Organization

```
docs/
├── assets/
│   ├── css/
│   │   └── kcsi.css          # Global brand stylesheet
│   └── images/
│       └── (SVG/PNG assets)
├── common.css                 # Legacy - to be merged with kcsi.css
├── index.html
├── cheatsheet.html
├── roadmap.html
├── teams.html
└── STYLE_GUIDE.md            # This file
```

---

## Maintenance Notes

**Solo Maintainer Budget:** 1-3 hours/week

- **Add new pages:** Copy `index.html` structure, update content
- **Update colors:** Modify CSS variables in `kcsi.css`, rebuild
- **Typography changes:** Update Google Fonts import, cascade through docs
- **Quick wins:** Use existing components, avoid custom one-offs

**Contact:** alessandro.middei@gmail.com
