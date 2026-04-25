/**
 * @license
 * SPDX-License-Identifier: Apache-2.0
 */

import React, { useEffect, useState } from 'react';
import { motion, AnimatePresence } from 'motion/react';
import { Search, Languages, ChevronUp, ChevronDown, MoreHorizontal } from 'lucide-react';
import { ArticleEntry, Category, Language } from './constants';
import uiText from './i18n/ui.json';

type Language = keyof typeof uiText;
type Token = keyof (typeof uiText)['en'];

const TOKENS = {
  brand: 'brand',
  navAi: 'nav.ai',
  navSocialTrends: 'nav.socialTrends',
  navMiscellaneous: 'nav.miscellaneous',
  search: 'actions.search',
  switchLanguage: 'actions.switchLanguage',
  contextualLinks: 'sidebar.contextualLinks',
  keyObservations: 'entry.keyObservations',
  footerCopyright: 'footer.copyright',
  footerPrivacyPolicy: 'footer.privacyPolicy',
  footerTermsOfService: 'footer.termsOfService',
  footerArchive: 'footer.archive',
  footerAbout: 'footer.about'
} as const satisfies Record<string, Token>;

function Navbar({
  t,
  activeCategory,
  onCategoryChange,
  onToggleLanguage
}: {
  t: (token: Token) => string;
  activeCategory: Category;
  onCategoryChange: (category: Category) => void;
  onToggleLanguage: () => void;
}) {
  return (
    <nav className="bg-paper-dim sticky top-0 w-full border-b border-border z-50">
      <div className="flex justify-between items-center max-w-container-max mx-auto px-8 h-20">
        {/* Brand */}
        <div className="font-serif text-3xl font-bold tracking-tight text-charcoal cursor-pointer">
          {t(TOKENS.brand)}
        </div>

        {/* Navigation Links */}
        <div className="hidden md:flex items-center space-x-8 font-serif italic text-lg opacity-80">
          <button
            type="button"
            onClick={() => onCategoryChange('ai')}
            className={activeCategory === 'ai'
              ? 'text-charcoal border-b-2 border-charcoal pb-1'
              : 'text-charcoal-muted hover:text-charcoal transition-all px-2 py-1 rounded-sm'}
          >
            {t(TOKENS.navAi)}
          </button>
          <button
            type="button"
            onClick={() => onCategoryChange('social-trends')}
            className={activeCategory === 'social-trends'
              ? 'text-charcoal border-b-2 border-charcoal pb-1'
              : 'text-charcoal-muted hover:text-charcoal transition-all px-2 py-1 rounded-sm'}
          >
            {t(TOKENS.navSocialTrends)}
          </button>
          <button
            type="button"
            onClick={() => onCategoryChange('miscellaneous')}
            className={activeCategory === 'miscellaneous'
              ? 'text-charcoal border-b-2 border-charcoal pb-1'
              : 'text-charcoal-muted hover:text-charcoal transition-all px-2 py-1 rounded-sm'}
          >
            {t(TOKENS.navMiscellaneous)}
          </button>
        </div>

        {/* Actions */}
        <div className="flex items-center space-x-4">
          <button
            aria-label={t(TOKENS.search)}
            className="hover:bg-black/5 transition-colors p-2 rounded-full flex items-center justify-center"
          >
            <Search size={20} className="text-charcoal" />
          </button>
          <button
            onClick={onToggleLanguage}
            aria-label={t(TOKENS.switchLanguage)}
            className="hover:bg-black/5 transition-colors px-3 py-2 rounded-full flex items-center justify-center font-sans text-sm text-charcoal"
          >
            <Languages size={18} />
          </button>
        </div>
      </div>
    </nav>
  );
}

function Sidebar({
  article,
  t
}: {
  article: ArticleEntry;
  t: (token: Token) => string;
}) {
  return (
    <aside className="lg:col-span-4 lg:col-start-9 bg-paper-dim border border-border p-6 h-fit mt-2 lg:mt-0">
      <h5 className="font-sans text-xs font-semibold uppercase tracking-widest mb-4 border-b border-border pb-2 opacity-60">
        {t(TOKENS.contextualLinks)}
      </h5>
      <ul className="space-y-4">
        {article.links.map((link, idx) => (
          <li key={idx}>
            <a className="block group/link" href={link.url}>
              <span className="font-sans text-sm text-charcoal group-hover/link:underline decoration-border underline-offset-4">
                {link.title}
              </span>
              <span className="block font-mono text-xs text-charcoal-muted mt-1 opacity-70">
                {link.domain}
              </span>
            </a>
          </li>
        ))}
      </ul>
    </aside>
  );
}

interface EntryProps {
  key?: React.Key;
  article: ArticleEntry;
  isExpanded: boolean;
  onToggle: () => void;
  t: (token: Token) => string;
}

function Entry({
  article,
  isExpanded,
  onToggle,
  t
}: EntryProps) {
  return (
    <article className={`group ${!isExpanded ? 'opacity-70 hover:opacity-100 transition-opacity' : ''}`}>
      <header
        className="flex items-end justify-between border-b border-border pb-2 mb-4 cursor-pointer group-hover:border-charcoal/30 transition-colors"
        onClick={onToggle}
      >
        <h2 className="font-serif text-3xl font-medium text-charcoal">{article.date}</h2>
        <div className="text-charcoal-muted group-hover:text-charcoal transition-colors mb-1">
          {isExpanded ? <ChevronUp size={20} /> : <ChevronDown size={20} />}
        </div>
      </header>

      <AnimatePresence>
        {isExpanded && (
          <motion.div
            initial={{ height: 0, opacity: 0 }}
            animate={{ height: 'auto', opacity: 1 }}
            exit={{ height: 0, opacity: 0 }}
            transition={{ duration: 0.4, ease: [0.23, 1, 0.32, 1] }}
            className="overflow-hidden"
          >
            <div className="grid grid-cols-1 lg:grid-cols-12 gap-8 py-4">
              <div className="lg:col-span-8 space-y-6">
                <h3 className="font-serif text-2xl font-medium">{article.title}</h3>
                <div className="markdown-body">
                  <p className="font-sans text-lg text-charcoal-muted leading-relaxed">
                    {article.content}
                  </p>
                </div>

                <div className="mt-8">
                  <h4 className="font-sans text-xs font-semibold uppercase tracking-widest mb-4 opacity-60">
                    {t(TOKENS.keyObservations)}
                  </h4>
                  <ul className="list-none space-y-3 pl-0 border-l border-border ml-2">
                    {article.observations.map((obs, idx) => (
                      <li
                        key={idx}
                        className="relative pl-6 before:absolute before:left-0 before:top-[12px] before:w-3 before:h-[1px] before:bg-border font-sans text-charcoal-muted"
                      >
                        {obs}
                      </li>
                    ))}
                  </ul>
                </div>

                <blockquote className="border-l-2 border-charcoal pl-6 py-2 my-8 italic font-serif text-2xl text-charcoal opacity-90 leading-relaxed">
                  "{article.quote}"
                </blockquote>
              </div>

              <Sidebar article={article} t={t} />
            </div>
          </motion.div>
        )}
      </AnimatePresence>
    </article>
  );
}

function Footer({ t }: { t: (token: Token) => string }) {
  return (
    <footer className="bg-paper-dim w-full border-t border-border mt-24">
      <div className="max-w-container-max mx-auto px-8 py-16 flex flex-col md:flex-row justify-between items-center gap-8">
        <div className="font-serif text-sm opacity-60">
          {t(TOKENS.footerCopyright)}
        </div>
        <div className="flex flex-wrap justify-center gap-6 font-serif text-sm opacity-60">
          <a className="hover:text-charcoal hover:underline transition-colors" href="#">{t(TOKENS.footerPrivacyPolicy)}</a>
          <a className="hover:text-charcoal hover:underline transition-colors" href="#">{t(TOKENS.footerTermsOfService)}</a>
          <a className="hover:text-charcoal hover:underline transition-colors" href="#">{t(TOKENS.footerArchive)}</a>
          <a className="hover:text-charcoal hover:underline transition-colors" href="#">{t(TOKENS.footerAbout)}</a>
        </div>
      </div>
    </footer>
  );
}

export default function App() {
  const [activeCategory, setActiveCategory] = useState<Category>('ai');
  const [expandedId, setExpandedId] = useState<string | null>(null);
  const [language, setLanguage] = useState<Language>('en');
  const [articles, setArticles] = useState<ArticleEntry[]>([]);
  const t = (token: Token) => uiText[language][token];

  useEffect(() => {
    let cancelled = false;
    async function loadArticles() {
      try {
        const response = await fetch(
          `/apis/v1/news?category=${encodeURIComponent(activeCategory)}&language=${encodeURIComponent(language)}`
        );
        if (!response.ok) {
          if (!cancelled) {
            setArticles([]);
          }
          return;
        }
        const payload = await response.json() as { articles?: ArticleEntry[] };
        if (!cancelled) {
          setArticles(Array.isArray(payload.articles) ? payload.articles : []);
        }
      } catch {
        if (!cancelled) {
          setArticles([]);
        }
      }
    }

    loadArticles();
    return () => {
      cancelled = true;
    };
  }, [activeCategory, language]);

  useEffect(() => {
    setExpandedId(articles[0]?.id ?? null);
  }, [articles]);

  return (
    <div className="min-h-screen flex flex-col selection:bg-accent/20">
      <Navbar
        t={t}
        activeCategory={activeCategory}
        onCategoryChange={setActiveCategory}
        onToggleLanguage={() => setLanguage((prev) => (prev === 'en' ? 'zh' : 'en'))}
      />

      <main className="flex-grow max-w-container-max mx-auto w-full px-8 py-16 space-y-20">
        {articles.map((article) => (
          <Entry
            key={article.id}
            article={article}
            isExpanded={expandedId === article.id}
            onToggle={() => setExpandedId(expandedId === article.id ? null : article.id)}
            t={t}
          />
        ))}

        {articles.length > 0 && (
          <div className="flex justify-center items-center py-12 opacity-30">
            <MoreHorizontal size={32} className="animate-pulse" />
          </div>
        )}
      </main>

      <Footer t={t} />
    </div>
  );
}
