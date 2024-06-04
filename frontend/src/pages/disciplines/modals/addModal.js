import { useForm, Controller } from "react-hook-form";
import styles from "../../shared/style/modal.module.css";
import Button from "../../../components/button";
import { useAddDisciplinesMutation } from "../api/addDisciplinesMutation";
import { useGroupsQuery } from "../../groups/api/getGroupsQuery";
import { useTeachersQuery } from "../../teachers/api/getTeachersQuery";
import { Multiselect } from "multiselect-react-dropdown";

const AddDisciplinesModal = ({ isOpen, toggleModal }) => {
  const form = useForm();
  const {
    reset,
    register,
    handleSubmit,
    control,
    formState: { errors },
  } = form;

  const groupsQuery = useGroupsQuery();
  const teachersQuery = useTeachersQuery();
  const addDisciplinesMutation = useAddDisciplinesMutation();

  const onSubmit = (data) => {
    addDisciplinesMutation.mutateAsync(data);
    toggleModal();
    reset();
  };

  const options = groupsQuery.data?.map((el) => `${el.groupId} ${el.id}`);

  const multiSelectStyles = {
    multiselectContainer: {
      backgroundColor: "white",
    },
    inputField: {
      margin: "5px",
    },
    chips: {
      background: "#8090bc",
    },
    optionContainer: {
      maxHeight: "160px",
      overflowY: "scroll",
    },
  };

  if (!isOpen) {
    return null;
  }

  return (
    <div className={styles["backdrop"]}>
      <div className={styles["modal"]}>
        <div className={styles["modal-header"]}>
          <h2>Введите данные</h2>
          <Button onClick={toggleModal} styleFeature="close">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              x="0px"
              y="0px"
              width="20"
              height="20"
              viewBox="0 0 30 30"
            >
              <path d="M 7 4 C 6.744125 4 6.4879687 4.0974687 6.2929688 4.2929688 L 4.2929688 6.2929688 C 3.9019687 6.6839688 3.9019687 7.3170313 4.2929688 7.7070312 L 11.585938 15 L 4.2929688 22.292969 C 3.9019687 22.683969 3.9019687 23.317031 4.2929688 23.707031 L 6.2929688 25.707031 C 6.6839688 26.098031 7.3170313 26.098031 7.7070312 25.707031 L 15 18.414062 L 22.292969 25.707031 C 22.682969 26.098031 23.317031 26.098031 23.707031 25.707031 L 25.707031 23.707031 C 26.098031 23.316031 26.098031 22.682969 25.707031 22.292969 L 18.414062 15 L 25.707031 7.7070312 C 26.098031 7.3170312 26.098031 6.6829688 25.707031 6.2929688 L 23.707031 4.2929688 C 23.316031 3.9019687 22.682969 3.9019687 22.292969 4.2929688 L 15 11.585938 L 7.7070312 4.2929688 C 7.5115312 4.0974687 7.255875 4 7 4 z"></path>
            </svg>
          </Button>
        </div>
        <form
          onSubmit={handleSubmit(onSubmit)}
          className={styles["groups-form"]}
        >
          <label>Введите id дисциплины</label>
          <input
            {...register("disciplinesId", { required: true })}
            placeholder="abc31"
            className={styles["input"]}
          />
          {errors.id && (
            <span className={styles.error}>Это поле обязательно</span>
          )}
          <label>Введите имя дисциплины</label>
          <input
            {...register("name", { required: true })}
            placeholder="Ин.яз"
            className={styles["input"]}
          />
          <label>Введите количество часов в неделю</label>
          <input
            {...register("hours", { required: true })}
            placeholder="20"
            className={styles["input"]}
          />
          <label>Выберете Преподователя</label>
          <select
            {...register("teachers", { required: true })}
            className={styles["input"]}
          >
            {teachersQuery.data.map((el, index) => (
              <option
                key={index}
                value={`${el.id}`}
              >{`Преподователь: ${el.name}`}</option>
            ))}
          </select>
          <label>Выберете группу(ы)</label>
          <Controller
            control={control}
            name="relatedGroupsId"
            render={({ field: { value, onChange } }) => (
              <Multiselect
                placeholder="314 GTH-JDO-NSK"
                options={options}
                isObject={false}
                showCheckbox={true}
                hidePlaceholder={true}
                closeOnSelect={false}
                onSelect={onChange}
                onRemove={onChange}
                selectedValues={value}
                showArrow={true}
                style={multiSelectStyles}
              />
            )}
          />
          <Button type="submit">Отправить</Button>
        </form>
      </div>
    </div>
  );
};

export default AddDisciplinesModal;
