const APP_NAME = 'digital-democracy-web';
// Keep in sync with package.json "version".
const APP_VERSION = '0.0.1';

export function appInfo(): { name: string; version: string } {
  return { name: APP_NAME, version: APP_VERSION };
}
