import { test, expect } from '@playwright/test';
import {Eyes, Target} from '@applitools/eyes-playwright';



test('basic visual test', async ({ page }) => {
    const eyes = new Eyes();
    eyes.setApiKey(String(process.env.APPLITOOLS_API_KEY!));
    await eyes.open(page, 'Personal Budget App', 'Basic Visual Test', {width: 800, height: 600});
    await page.goto('localhost:3000/');
    await eyes.check('Welcome the Personal Budget Application!', Target.window());
    await eyes.close();
});

test('sign in works', async ({ page }) => {
    await page.goto('localhost:3000/sign-in');
    await page.fill('input[name="email"]', 'test@test.com');
    await page.fill('input[name="password"]', process.env.TEST_PASSWORD!);
    await page.click('button[type="submit"]');
    await expect(page).
    toHaveURL('http://localhost:3000/dashboard', { timeout: 10000 });
});

test('sign in fails', async ({ page }) => {
    await page.goto('localhost:3000/sign-in');
    await page.fill('input[name="email"]', 'test@test.com');
    await page.fill('input[name="password"]', 'wrongpassword');
    await page.click('button[type="submit"]');
    await expect(page.locator('#confirmation-message')).toHaveText('username and/or password incorrect');
});
