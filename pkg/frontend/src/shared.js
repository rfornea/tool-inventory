export const BASE_URI = process.env.VUE_APP_HOST;

const getImageUrl = id => {
  return BASE_URI + "/assets/images/" + id + ".png";
};

export default getImageUrl;
