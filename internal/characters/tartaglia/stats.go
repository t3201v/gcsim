package tartaglia

var (
	attack = [][]float64{
		{
			0.4128,
			0.4464,
			0.48,
			0.528,
			0.5616,
			0.6,
			0.6528,
			0.7056,
			0.7584,
			0.816,
			0.8736,
			0.9312,
			0.9888,
			1.0464,
			1.104,
		},
		{
			0.4627,
			0.5003,
			0.538,
			0.5918,
			0.6295,
			0.6725,
			0.7317,
			0.7909,
			0.85,
			0.9146,
			0.9792,
			1.0437,
			1.1083,
			1.1728,
			1.2374,
		},
		{
			0.5538,
			0.5989,
			0.644,
			0.7084,
			0.7535,
			0.805,
			0.8758,
			0.9467,
			1.0175,
			1.0948,
			1.1721,
			1.2494,
			1.3266,
			1.4039,
			1.4812,
		},
		{
			0.5702,
			0.6166,
			0.663,
			0.7293,
			0.7757,
			0.8288,
			0.9017,
			0.9746,
			1.0475,
			1.1271,
			1.2067,
			1.2862,
			1.3658,
			1.4453,
			1.5249,
		},
		{
			0.6089,
			0.6584,
			0.708,
			0.7788,
			0.8284,
			0.885,
			0.9629,
			1.0408,
			1.1186,
			1.2036,
			1.2886,
			1.3735,
			1.4585,
			1.5434,
			1.6284,
		},
		{
			0.7276,
			0.7868,
			0.846,
			0.9306,
			0.9898,
			1.0575,
			1.1506,
			1.2436,
			1.3367,
			1.4382,
			1.5397,
			1.6412,
			1.7428,
			1.8443,
			1.9458,
		},
	}
	aim = []float64{
		1.24,
		1.333,
		1.426,
		1.55,
		1.643,
		1.736,
		1.86,
		1.984,
		2.108,
		2.232,
		2.356,
		2.48,
		2.635,
		2.79,
		2.945,
	}
	rtFlash = [][]float64{
		{
			0.124,
			0.1333,
			0.1426,
			0.155,
			0.1643,
			0.1736,
			0.186,
			0.1984,
			0.2108,
			0.2232,
			0.2356,
			0.248,
			0.2635,
			0.279,
			0.2945,
		},
		{
			0.124,
			0.1333,
			0.1426,
			0.155,
			0.1643,
			0.1736,
			0.186,
			0.1984,
			0.2108,
			0.2232,
			0.2356,
			0.248,
			0.2635,
			0.279,
			0.2945,
		},
		{
			0.124,
			0.1333,
			0.1426,
			0.155,
			0.1643,
			0.1736,
			0.186,
			0.1984,
			0.2108,
			0.2232,
			0.2356,
			0.248,
			0.2635,
			0.279,
			0.2945,
		},
	}
	rtBurst = []float64{
		0.62,
		0.6665,
		0.713,
		0.775,
		0.8215,
		0.868,
		0.93,
		0.992,
		1.054,
		1.116,
		1.178,
		1.24,
		1.3175,
		1.395,
		1.4725,
	}
	skill = []float64{
		0.72,
		0.774,
		0.828,
		0.9,
		0.954,
		1.008,
		1.08,
		1.152,
		1.224,
		1.296,
		1.368,
		1.44,
		1.53,
		1.62,
		1.71,
	}
	eAttack = [][][]float64{
		{ //1
			{
				0.3887,
				0.4204,
				0.452,
				0.4972,
				0.5288,
				0.565,
				0.6147,
				0.6644,
				0.7142,
				0.7684,
				0.8226,
				0.8769,
				0.9311,
				0.9854,
				1.0396,
			},
		},
		{ //2
			{
				0.4162,
				0.4501,
				0.484,
				0.5324,
				0.5663,
				0.605,
				0.6582,
				0.7115,
				0.7647,
				0.8228,
				0.8809,
				0.939,
				0.997,
				1.0551,
				1.1132,
			},
		},
		{ //3
			{
				0.5633,
				0.6092,
				0.655,
				0.7205,
				0.7664,
				0.8188,
				0.8908,
				0.9629,
				1.0349,
				1.1135,
				1.1921,
				1.2707,
				1.3493,
				1.4279,
				1.5065,
			},
		},
		{ //4
			{
				0.5994,
				0.6482,
				0.697,
				0.7667,
				0.8155,
				0.8713,
				0.9479,
				1.0246,
				1.1013,
				1.1849,
				1.2685,
				1.3522,
				1.4358,
				1.5195,
				1.6031,
			},
		},
		{ //5
			{
				0.553,
				0.598,
				0.643,
				0.7073,
				0.7523,
				0.8038,
				0.8745,
				0.9452,
				1.0159,
				1.0931,
				1.1703,
				1.2474,
				1.3246,
				1.4017,
				1.4789,
			},
		},
		{ //6
			{
				0.3543,
				0.3832,
				0.412,
				0.4532,
				0.482,
				0.515,
				0.5603,
				0.6056,
				0.651,
				0.7004,
				0.7498,
				0.7993,
				0.8487,
				0.8982,
				0.9476,
			},
			{
				0.3767,
				0.4073,
				0.438,
				0.4818,
				0.5125,
				0.5475,
				0.5957,
				0.6439,
				0.692,
				0.7446,
				0.7972,
				0.8497,
				0.9023,
				0.9548,
				1.0074,
			},
		},
	}
	eCharge = [][]float64{
		{
			0.602,
			0.651,
			0.70,
			0.77,
			0.819,
			0.875,
			0.952,
			1.029,
			1.106,
			1.19,
			1.274,
			1.358,
			1.442,
			1.526,
			1.61,
		},
		{
			0.7198,
			0.7784,
			0.837,
			0.9207,
			0.9793,
			1.0462,
			1.1383,
			1.2304,
			1.3225,
			1.4229,
			1.5233,
			1.6238,
			1.7242,
			1.8247,
			1.9251,
		},
	}
	rtSlash = []float64{
		0.602,
		0.651,
		0.7,
		0.77,
		0.819,
		0.875,
		0.952,
		1.029,
		1.106,
		1.19,
		1.274,
		1.358,
		1.442,
		1.526,
		1.61,
	}
	meleeBurst = []float64{
		4.64,
		4.988,
		5.336,
		5.8,
		6.148,
		6.496,
		6.96,
		7.424,
		7.888,
		8.352,
		8.816,
		9.28,
		9.86,
		10.44,
		11.02,
	}
	burst = []float64{
		3.784,
		4.0678,
		4.3516,
		4.73,
		5.0138,
		5.2976,
		5.676,
		6.0544,
		6.4328,
		6.8112,
		7.1896,
		7.568,
		8.041,
		8.514,
		8.987,
	}
	rtBlast = []float64{
		1.2,
		1.29,
		1.38,
		1.5,
		1.59,
		1.68,
		1.8,
		1.92,
		2.04,
		2.16,
		2.28,
		2.4,
		2.55,
		2.7,
		2.85,
	}
)
