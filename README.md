# go-silver-gold-tracker

Lists gold and silver price in USD on request.

## Output

```shell
20:18 Ag: 20.94$; Au: 1819.93$
<time of last request> Ag: <silver price in usd per ounce>; Au: <gold price in usd per ounce>
```

## API

This uses the [gold.de](https://www.gold.de/chartgenerator/) public api. It's free, it's simple, and it works.

<https://api.edelmetalle.de/public.json>

```json
{
  "gold_usd": 1819.7,
  "gold_eur": 1726.09,
  "silber_usd": 20.939,
  "silber_eur": 19.8615,
  "platin_usd": 863.465,
  "platin_eur": 819.055,
  "palladium_usd": 1148.85,
  "palladium_eur": 1089.76,
  "timestamp": 1696529768,
  "wechselkurs_usd_eur": 1.0538818076477405
}
```