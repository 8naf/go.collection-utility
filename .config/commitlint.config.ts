import commitTypes from './commit-types.json'

const types = commitTypes.map(({ type, emoji }) => {
	return `${emoji} ${type}`
})

export default {
	extends: ['@commitlint/config-conventional'],
	parserPreset: {
		parserOpts: {
			headerPattern: /^(:\w+: \w+)(?:\((.*)\))?!?: (.*)$/,
			headerCorrespondence: ['type', 'scope', 'subject'],
		},
	},
	rules: {
		'type-enum': [2, 'always', types],
	},
}
