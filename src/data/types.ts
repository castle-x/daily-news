export type Language = 'en' | 'zh';

export type Category = 'ai' | 'social-trends' | 'miscellaneous';

export interface LocalizedText {
  en: string;
  zh: string;
}

export interface BilingualLink {
  title: LocalizedText;
  url: string;
  domain: string;
}

export interface BilingualArticleRecord {
  id: string;
  date: string;
  category: Category;
  title: LocalizedText;
  summary: LocalizedText;
  observations: LocalizedText[];
  quote: LocalizedText;
  links: BilingualLink[];
}

export interface ArticleLink {
  title: string;
  url: string;
  domain: string;
}

export interface ArticleEntry {
  id: string;
  date: string;
  title: string;
  content: string;
  observations: string[];
  quote: string;
  links: ArticleLink[];
}
