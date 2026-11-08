import { describe, expect, it } from 'vitest';
import { appInfo } from './app-info';

describe('appInfo', () => {
  it('returns name and version', () => {
    const info = appInfo();
    expect(info.name).toBe('digital-democracy-web');
    expect(info.version).toBe('0.0.1');
  });
});
