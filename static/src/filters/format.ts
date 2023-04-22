import {DateTime, Duration} from "luxon"

const floatFormatter: Intl.NumberFormat = new Intl.NumberFormat('local', {
        minimumFractionDigits: 0,
        maximumFractionDigits: 3,
    }),
    intFormatter: Intl.NumberFormat = new Intl.NumberFormat()

const kB = 1024,
    mB = kB * 1024,
    gB = mB * 1024,
    tB = gB * 1024

const filters = {
    formatFloat(value: number): string {
        if (value === undefined || value === null || isNaN(value))
            return '-'

        return floatFormatter.format(value)
    },

    formatInt(value: number): string {
        if (value === undefined || value === null || isNaN(value))
            return '-'

        return intFormatter.format(Math.round(value))
    },

    formatBytes(value: number): string {
        if (value === undefined || value === null || isNaN(value))
            return '-'

        if (value > tB)
            return floatFormatter.format(value / tB) + 'tB'

        if (value > gB)
            return floatFormatter.format(value / gB) + 'gB'

        if (value > mB)
            return floatFormatter.format(value / mB) + 'mB'

        if (value > kB)
            return floatFormatter.format(value / kB) + 'kB'

        return floatFormatter.format(value) + 'B'
    },

    formatDateTime(dt: number | string): string {
        switch (typeof dt) {
            case "string":
                return DateTime.fromISO(dt).setLocale('en-GB').toLocaleString(DateTime.DATETIME_MED_WITH_SECONDS)
            case "number":
                return DateTime.fromMillis(dt).setLocale('en-GB').toLocaleString(DateTime.DATETIME_MED_WITH_SECONDS)
        }
    },

    formatDate(dt: number | string): string {
        switch (typeof dt) {
            case "string":
                return DateTime.fromISO(dt).setLocale('en-GB').toLocaleString(DateTime.DATE_MED)
            case "number":
                return DateTime.fromMillis(dt).setLocale('en-GB').toLocaleString(DateTime.DATE_MED)
        }
    },

    formatDuration(ms: number): string {
        if (ms === undefined || ms === null || isNaN(ms)) {
            return '—'
        }

        if (ms == 0) {
            return '0ms'
        }

        const durParts = Duration.fromMillis(ms).shiftTo('hours', 'minutes', 'seconds', 'milliseconds').toObject()

        const res: string[] = []

        if (durParts.hours)
            res.push(`${durParts.hours}h`)

        if (durParts.minutes)
            res.push(`${durParts.minutes}m`)

        if (durParts.seconds)
            res.push(`${durParts.seconds}s`)

        if (durParts.milliseconds)
            res.push(`${floatFormatter.format(durParts.milliseconds)}ms`)

        return res.join(' ')
    }
}


export default filters
